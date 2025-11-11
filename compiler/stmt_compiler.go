package compiler

import (
	"fmt"
	"os"
	"pepper/lexer"
	"pepper/parser"
	"pepper/runtime"
)

func (c *Compiler) compileStmt(stmt parser.Statement, isExprContext bool) {
	if stmt == nil {
		return
	}

	switch node := stmt.(type) {
	case *parser.ExpressionStatement:
		c.compileExpr(node.Expression)
	case *parser.DimStatement:
		c.compileDimStatement(node)
	case *parser.ReturnStatement:
		c.compileExpr(node.ReturnValue)
		c.emit(runtime.OpReturn)
	case *parser.BlockStatement:
		last := len(node.Statements) - 1
		for i, s := range node.Statements {
			c.compileStmt(s, isExprContext && (i == last))
		}
	case *parser.LoopStatement:
		c.compileLoopStatement(node)
	case *parser.RepeatStatement:
		c.compileRepeatStatement(node)
	case *parser.BreakStatement:
		c.compileBreakStatement(node)
	case *parser.ContinueStatement:
		c.compileContinueStatement(node)

	case *parser.IncludeStatement:
		program, ok := c.includedPrograms[node.Filename]
		if !ok {
			data, err := os.ReadFile(node.Filename)
			if err != nil {
				panic(err)
			}
			l := lexer.New(string(data))
			p := parser.New(l)
			program = p.ParseProgram()
			if len(p.Errors()) != 0 {
				for _, msg := range p.Errors() {
					fmt.Println(msg)
				}
				panic("Parser errors")
			}
			c.includedPrograms[node.Filename] = program
		}
		for _, stmt := range program.Statements {
			c.compileStmt(stmt, false)
		}

	case *parser.MatchExpression:
		c.compileMatchExpression(node)

	default:
		token := stmt.GetToken()
		panic(fmt.Sprintf("line %d:%d: Unsupported statement type: %T", token.Line, token.Column, stmt))
	}
}

func (c *Compiler) compileDimStatement(node *parser.DimStatement) {
	if node.Value == nil {
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.NIL, Value: nil}) // Push nil
	} else {
		c.compileExpr(node.Value)
	}
	symbol := c.symbolTable.DefineVar(node.Name.Value)
	c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, Value: symbol.Name})
	if symbol.Scope == GlobalScope {
		c.emit(runtime.OpStoreGlobal)
	} else {
		c.emit(runtime.OpStoreLocal)
	}
}

func (c *Compiler) compileLoopStatement(node *parser.LoopStatement) {
	loopStart := len(c.instructions)
	c.pushLoopContext(loopStart)

	// Attempt to compile for combined compare-and-jump optimization
	if infix, ok := node.Condition.(*parser.InfixExpression); ok {
		jumpOp := c.getJumpOpForInfix(infix.Operator)
		if jumpOp != 0 {
			c.compileExpr(infix.Left)
			c.compileExpr(infix.Right)

			jmpIfFalsePos := c.emitWithPlaceholder(jumpOp)

			c.enterScope()
			c.compileStmt(node.Body, false)
			c.leaveScope()

			c.emit(runtime.OpJmp, runtime.VMDataObject{Type: runtime.INTGER, Value: int64(loopStart)})

			loopEnd := len(c.instructions)
			c.patchJump(jmpIfFalsePos)
			c.patchBreaks(loopEnd)
			c.popLoopContext()
			return
		}
	}

	// Fallback for single values or non-optimizable expressions
	c.compileExpr(node.Condition)
	jmpIfFalsePos := c.emitWithPlaceholder(runtime.OpJmpIfFalse)

	c.enterScope()
	c.compileStmt(node.Body, false)
	c.leaveScope()

	c.emit(runtime.OpJmp, runtime.VMDataObject{Type: runtime.INTGER, Value: int64(loopStart)})

	loopEnd := len(c.instructions)
	c.patchJump(jmpIfFalsePos)
	c.patchBreaks(loopEnd)
	c.popLoopContext()
}

func (c *Compiler) compileRepeatStatement(node *parser.RepeatStatement) {
	count, ok := node.Count.(*parser.IntegerLiteral)
	if !ok {
		token := node.GetToken()
		panic(fmt.Sprintf("line %d:%d: `repeat` currently only supports integer literals for the count.", token.Line, token.Column))
	}

	loopStart := len(c.instructions)
	c.pushLoopContext(loopStart)

	for i := int64(0); i < count.Value; i++ {
		c.enterScope()
		c.compileStmt(node.Body, false)
		c.leaveScope()
	}

	loopEnd := len(c.instructions)
	c.patchBreaks(loopEnd)
	c.popLoopContext()
}

func (c *Compiler) compileBreakStatement(node *parser.BreakStatement) {
	if len(c.loopContexts) == 0 {
		token := node.GetToken()
		panic(fmt.Sprintf("line %d:%d: 'break' outside of a loop", token.Line, token.Column))
	}
	patchPos := c.emitWithPlaceholder(runtime.OpJmp)
	currentLoop := c.loopContexts[len(c.loopContexts)-1]
	currentLoop.breakPatches = append(currentLoop.breakPatches, patchPos)
}

func (c *Compiler) compileContinueStatement(node *parser.ContinueStatement) {
	if len(c.loopContexts) == 0 {
		token := node.GetToken()
		panic(fmt.Sprintf("line %d:%d: 'continue' outside of a loop", token.Line, token.Column))
	}
	currentLoop := c.loopContexts[len(c.loopContexts)-1]
	c.emit(runtime.OpJmp, runtime.VMDataObject{Type: runtime.INTGER, Value: int64(currentLoop.startPos)})
}

func isReturnStatement(stmt parser.Statement) bool {
	_, ok := stmt.(*parser.ReturnStatement)
	return ok
}

package compiler

import (
	"fmt"
	"os"
	"pepper/lexer"
	"pepper/parser"
	"pepper/runtime"
	"reflect"
)

type SymbolScope string

const (
	GlobalScope SymbolScope = "GLOBAL"
	LocalScope  SymbolScope = "LOCAL"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
	Outer          *SymbolTable
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions}
	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
		return obj, ok
	}
	return obj, ok
}

type LoopContext struct {
	startPos     int
	breakPatches []int
}

type Compiler struct {
	instructions         []runtime.VMInstr
	loopContexts         []*LoopContext
	standardFunctionMaps map[string][]runtime.VMInstr
	symbolTable          *SymbolTable
}

func NewCompiler() *Compiler {
	symbolTable := NewSymbolTable()
	return &Compiler{
		loopContexts:         make([]*LoopContext, 0),
		standardFunctionMaps: make(map[string][]runtime.VMInstr),
		symbolTable:          symbolTable,
	}
}

func (c *Compiler) enterScope() {
	c.symbolTable = NewEnclosedSymbolTable(c.symbolTable)
}

func (c *Compiler) leaveScope() {
	c.symbolTable = c.symbolTable.Outer
}

func (c *Compiler) Compile(program *parser.Program, excludestd bool) []runtime.VMInstr {
	if !excludestd {
		c.defineStandardFunctions()
		for name, instrs := range c.standardFunctionMaps {
			c.instructions = append(c.instructions, runtime.VMInstr{Op: runtime.OpDefFunc, Oprand1: runtime.VMDataObject{Type: runtime.STRING, StringData: name}})
			c.instructions = append(c.instructions, runtime.VMInstr{Op: runtime.OpJmp, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: int64(len(c.instructions)) + 3}})
			c.instructions = append(c.instructions, instrs...)
			c.instructions = append(c.instructions, runtime.VMInstr{Op: runtime.OpReturn})
		}
	}

	for _, stmt := range program.Statements {
		c.compileStmt(stmt, false)
	}
	return c.instructions
}

func (c *Compiler) compileStmt(stmt parser.Statement, isExprContext bool) {
	if stmt == nil || (reflect.ValueOf(stmt).Kind() == reflect.Ptr && reflect.ValueOf(stmt).IsNil()) {
		return
	}

	switch node := stmt.(type) {
	case *parser.ExpressionStatement:
		c.compileExpr(node.Expression)
	case *parser.LetStatement:
		c.compileLetStatement(node)
	case *parser.DimStatement:
		c.compileDimStatement(node)
	case *parser.ReturnStatement:
		c.compileExpr(node.ReturnValue)
		c.emit(runtime.OpReturn)
	case *parser.BlockStatement:
		if len(node.Statements) == 0 {
			if isExprContext {
				c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.NIL})
			}
			return
		}
		last := len(node.Statements) - 1
		for i, s := range node.Statements {
			c.compileStmt(s, isExprContext && (i == last))
		}
	case *parser.LoopStatement:
		c.compileLoopStatement(node)
	case *parser.RepeatStatement:
		c.compileRepeatStatement(node)
	case *parser.BreakStatement:
		c.compileBreakStatement()
	case *parser.ContinueStatement:
		c.compileContinueStatement()

	case *parser.IncludeStatement:
		data, err := os.ReadFile(node.Filename)
		if err != nil {
			panic(err)
		}
		l := lexer.New(string(data))
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			for _, msg := range p.Errors() {
				fmt.Println(msg)
			}
			panic("Parser errors")
		}
		for _, stmt := range program.Statements {
			c.compileStmt(stmt, false)
		}

	default:
		panic(fmt.Sprintf("Unsupported statement type: %T", stmt))
	}
}

func (c *Compiler) compileExpr(expr parser.Expression) {
	switch node := expr.(type) {
	case *parser.InfixExpression:
		c.compileInfixExpr(node)
	case *parser.PrefixExpression:
		c.compilePrefixExpr(node)
	case *parser.IfExpression:
		c.compileIfExpression(node)
	case *parser.BlockExpression:
		c.compileBlockExpression(node)
	case *parser.FunctionLiteral:
		c.compileFunctionLiteral(node)
	case *parser.CallExpression:
		c.compileCallExpression(node)
	case *parser.AssignmentExpression:
		c.compileAssignmentExpression(node)
	case *parser.IndexExpression:
		c.compileIndexExpression(node)
	case *parser.MemberAccessExpression:
		c.compileMemberAccessExpression(node)
	case *parser.IntegerLiteral:
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.INTGER, IntData: node.Value})
	case *parser.RealLiteral:
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.REAL, FloatData: node.Value})
	case *parser.StringLiteral:
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: node.Value})
	case *parser.Boolean:
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.BOOLEAN, BoolData: node.Value})
	case *parser.NilLiteral:
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.NIL}) // Nil
	case *parser.Identifier:
		symbol, ok := c.symbolTable.Resolve(node.Value)
		if !ok {
			panic(fmt.Sprintf("undefined variable: %s", node.Value))
		}
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: symbol.Name})
		if symbol.Scope == GlobalScope {
			c.emit(runtime.OpLoadGlobal)
		} else {
			c.emit(runtime.OpLoadLocal)
		}
	case *parser.PackLiteral:
		c.compilePackLiteral(node)
	default:
		panic(fmt.Sprintf("Unsupported expression type: %T", expr))
	}
}

func (c *Compiler) compileBlockExpression(node *parser.BlockExpression) {
	if len(node.Statements) == 0 {
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.NIL}) // Nil for empty block
		return
	}
	last := len(node.Statements) - 1
	for i, s := range node.Statements {
		isLast := (i == last)
		c.compileStmt(s, isLast) // Only the last statement is in an expression context
	}
}

func (c *Compiler) compileLetStatement(node *parser.LetStatement) {
	c.compileExpr(node.Value)
	symbol := c.symbolTable.Define(node.Name.Value)
	c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: symbol.Name})
	if symbol.Scope == GlobalScope {
		c.emit(runtime.OpStoreGlobal)
	} else {
		c.emit(runtime.OpStoreLocal)
	}
}

func (c *Compiler) compileDimStatement(node *parser.DimStatement) {
	if node.Value == nil {
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.NIL}) // Push nil
	} else {
		c.compileExpr(node.Value)
	}
	symbol := c.symbolTable.Define(node.Name.Value)
	c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: symbol.Name})
	if symbol.Scope == GlobalScope {
		c.emit(runtime.OpStoreGlobal)
	} else {
		c.emit(runtime.OpStoreLocal)
	}
}

func (c *Compiler) compileInfixExpr(node *parser.InfixExpression) {
	c.compileExpr(node.Left)
	c.compileExpr(node.Right)

	switch node.Operator {
	case "+":
		c.emit(runtime.OpAdd)
	case "-":
		c.emit(runtime.OpSub)
	case "*":
		c.emit(runtime.OpMul)
	case "/":
		c.emit(runtime.OpDiv)
	case "%":
		c.emit(runtime.OpMod)
	case "==":
		c.emit(runtime.OpCmpEq)
	case "!=":
		c.emit(runtime.OpCmpNeq)
	case ">":
		c.emit(runtime.OpCmpGt)
	case "<":
		c.emit(runtime.OpCmpLt)
	case ">=":
		c.emit(runtime.OpCmpGte)
	case "<=":
		c.emit(runtime.OpCmpLte)
	case "and":
		c.emit(runtime.OpAnd)
	case "or":
		c.emit(runtime.OpOr)
	default:
		panic(fmt.Sprintf("Unknown infix operator: %s", node.Operator))
	}
}

func (c *Compiler) compilePrefixExpr(node *parser.PrefixExpression) {
	c.compileExpr(node.Right)
	switch node.Operator {
	case "not":
		c.emit(runtime.OpNot)
	case "-":
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.INTGER, IntData: 0})
		c.emit(runtime.OpSub)
	default:
		panic(fmt.Sprintf("Unknown prefix operator: %s", node.Operator))
	}
}

func (c *Compiler) compileIfExpression(node *parser.IfExpression) {
	c.compileExpr(node.Condition)
	jmpIfFalsePos := c.emitWithPlaceholder(runtime.OpJmpIfFalse)

	if len(node.Consequence.Statements) == 0 {
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.NIL})
	} else {
		c.compileStmt(node.Consequence, true)
	}

	jmpPos := c.emitWithPlaceholder(runtime.OpJmp)
	c.patchJump(jmpIfFalsePos)

	if node.Alternative == nil {
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.NIL})
	} else {
		c.compileStmt(node.Alternative, true)
	}

	c.patchJump(jmpPos)
}

func (c *Compiler) compileFunctionLiteral(node *parser.FunctionLiteral) {
	c.emit(runtime.OpDefFunc, runtime.VMDataObject{Type: runtime.STRING, StringData: node.Name.Value})
	jumpPos := c.emitWithPlaceholder(runtime.OpJmp)

	c.enterScope()

	for i := len(node.Parameters) - 1; i >= 0; i-- {
		pname := node.Parameters[i]
		c.symbolTable.Define(pname.Value)
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: pname.Value})
		c.emit(runtime.OpStoreLocal)
	}

	c.compileStmt(node.Body, true)
	if len(node.Body.Statements) == 0 || !isReturnStatement(node.Body.Statements[len(node.Body.Statements)-1]) {
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.NIL})
		c.emit(runtime.OpReturn)
	}

	c.leaveScope()
	c.patchJump(jumpPos)
}

func (c *Compiler) compileCallExpression(node *parser.CallExpression) {
	for _, arg := range node.Arguments {
		c.compileExpr(arg)
	}

	ident, ok := node.Function.(*parser.Identifier)
	if !ok {
		panic("Calling non-identifier function is not supported yet")
	}

	c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: ident.Value})
	c.emit(runtime.OpCall)
}

func (c *Compiler) compileAssignmentExpression(node *parser.AssignmentExpression) {
	if ident, ok := node.Left.(*parser.Identifier); ok {
		c.compileExpr(node.Value)
		symbol, ok := c.symbolTable.Resolve(ident.Value)
		if !ok {
			panic(fmt.Sprintf("undefined variable: %s", ident.Value))
		}
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: symbol.Name})
		if symbol.Scope == GlobalScope {
			c.emit(runtime.OpStoreGlobal)
		} else {
			c.emit(runtime.OpStoreLocal)
		}
	} else if indexExpr, ok := node.Left.(*parser.IndexExpression); ok {
		c.compileExpr(indexExpr.Left)
		c.compileExpr(indexExpr.Index)
		c.compileExpr(node.Value)
		c.emit(runtime.OpSetIndex)

		if ident, ok := indexExpr.Left.(*parser.Identifier); ok {
			symbol, ok := c.symbolTable.Resolve(ident.Value)
			if !ok {
				panic(fmt.Sprintf("undefined variable: %s", ident.Value))
			}
			c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: symbol.Name})
			if symbol.Scope == GlobalScope {
				c.emit(runtime.OpStoreGlobal)
			} else {
				c.emit(runtime.OpStoreLocal)
			}
		} else {
			c.emit(runtime.OpPop)
		}
	} else if memberAccessExpr, ok := node.Left.(*parser.MemberAccessExpression); ok {
		c.compileExpr(memberAccessExpr.Object)
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: memberAccessExpr.Member.Value})
		c.compileExpr(node.Value)
		c.emit(runtime.OpSetIndex)

		if ident, ok := memberAccessExpr.Object.(*parser.Identifier); ok {
			symbol, ok := c.symbolTable.Resolve(ident.Value)
			if !ok {
				panic(fmt.Sprintf("undefined variable: %s", ident.Value))
			}
			c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: symbol.Name})
			if symbol.Scope == GlobalScope {
				c.emit(runtime.OpStoreGlobal)
			} else {
				c.emit(runtime.OpStoreLocal)
			}
		} else {
			c.emit(runtime.OpPop)
		}
	} else {
		panic(fmt.Sprintf("Assignment to this expression type is not supported: %T", node.Left))
	}
}

func (c *Compiler) compileLoopStatement(node *parser.LoopStatement) {
	loopStart := len(c.instructions)
	c.pushLoopContext(loopStart)

	c.compileExpr(node.Condition)
	jmpIfFalsePos := c.emitWithPlaceholder(runtime.OpJmpIfFalse)

	c.compileStmt(node.Body, false)

	c.emit(runtime.OpJmp, runtime.VMDataObject{Type: runtime.INTGER, IntData: int64(loopStart)})

	loopEnd := len(c.instructions)
	c.patchJump(jmpIfFalsePos)
	c.patchBreaks(loopEnd)
	c.popLoopContext()
}

func (c *Compiler) compileRepeatStatement(node *parser.RepeatStatement) {
	count, ok := node.Count.(*parser.IntegerLiteral)
	if !ok {
		panic("`repeat` currently only supports integer literals for the count.")
	}

	loopStart := len(c.instructions)
	c.pushLoopContext(loopStart)

	for i := int64(0); i < count.Value; i++ {
		c.compileStmt(node.Body, false)
	}

	loopEnd := len(c.instructions)
	c.patchBreaks(loopEnd)
	c.popLoopContext()
}

func (c *Compiler) compileBreakStatement() {
	if len(c.loopContexts) == 0 {
		panic("'break' outside of a loop")
	}
	patchPos := c.emitWithPlaceholder(runtime.OpJmp)
	currentLoop := c.loopContexts[len(c.loopContexts)-1]
	currentLoop.breakPatches = append(currentLoop.breakPatches, patchPos)
}

func (c *Compiler) compileContinueStatement() {
	if len(c.loopContexts) == 0 {
		panic("'continue' outside of a loop")
	}
	currentLoop := c.loopContexts[len(c.loopContexts)-1]
	c.emit(runtime.OpJmp, runtime.VMDataObject{Type: runtime.INTGER, IntData: int64(currentLoop.startPos)})
}

func (c *Compiler) compilePackLiteral(node *parser.PackLiteral) {
	c.emit(runtime.OpMakePack)

	for key, value := range node.Pairs {
		c.compileExpr(key)
		c.compileExpr(value)
		c.emit(runtime.OpSetIndex)
	}
}

func (c *Compiler) compileIndexExpression(node *parser.IndexExpression) {
	c.compileExpr(node.Left)
	c.compileExpr(node.Index)
	c.emit(runtime.OpIndex)
}

func (c *Compiler) compileMemberAccessExpression(node *parser.MemberAccessExpression) {
	c.compileExpr(node.Object)
	c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, StringData: node.Member.Value})
	c.emit(runtime.OpIndex)
}

func (c *Compiler) emit(op runtime.VMOp, operands ...runtime.VMDataObject) {
	instr := runtime.VMInstr{Op: op}
	if len(operands) > 0 {
		instr.Oprand1 = operands[0]
	}
	c.instructions = append(c.instructions, instr)
}

func (c *Compiler) emitWithPlaceholder(op runtime.VMOp) int {
	instr := runtime.VMInstr{Op: op, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: -1}}
	c.instructions = append(c.instructions, instr)
	return len(c.instructions) - 1
}

func (c *Compiler) patchJump(pos int) {
	jumpTo := len(c.instructions)
	c.instructions[pos].Oprand1.IntData = int64(jumpTo)
}

func (c *Compiler) pushLoopContext(start int) {
	context := &LoopContext{startPos: start, breakPatches: make([]int, 0)}
	c.loopContexts = append(c.loopContexts, context)
}

func (c *Compiler) popLoopContext() {
	c.loopContexts = c.loopContexts[:len(c.loopContexts)-1]
}

func (c *Compiler) patchBreaks(loopEnd int) {
	if len(c.loopContexts) == 0 {
		return
	}
	lastCtx := len(c.loopContexts) - 1
	for _, patchPos := range c.loopContexts[lastCtx].breakPatches {
		c.instructions[patchPos].Oprand1.IntData = int64(loopEnd)
	}
}

func isReturnStatement(stmt parser.Statement) bool {
	_, ok := stmt.(*parser.ReturnStatement)
	return ok
}

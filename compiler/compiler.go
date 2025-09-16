package compiler

import (
	"fmt"
	"pepper/parser"
	"pepper/vm"
	"reflect"
)

// LoopContext holds information about a loop being compiled,
// used to handle break and continue statements.
type LoopContext struct {
	startPos     int
	breakPatches []int
}

type Compiler struct {
	instructions         []vm.VMInstr
	loopContexts         []*LoopContext
	standardFunctionMaps map[string][]vm.VMInstr
}

func NewCompiler() *Compiler {
	return &Compiler{
		loopContexts:         make([]*LoopContext, 0),
		standardFunctionMaps: make(map[string][]vm.VMInstr),
	}
}

func (c *Compiler) Compile(program *parser.Program) []vm.VMInstr {
	c.defineStandardFunctions()
	for name, instrs := range c.standardFunctionMaps {
		c.instructions = append(c.instructions, vm.VMInstr{Op: vm.OpDefFunc, Oprand1: vm.VMDataObject{Type: vm.STRING, StringData: name}})
		c.instructions = append(c.instructions, instrs...)
		c.instructions = append(c.instructions, vm.VMInstr{Op: vm.OpReturn})
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
		c.emit(vm.OpReturn)
	case *parser.BlockStatement:
		for _, s := range node.Statements {
			c.compileStmt(s, false) // Statements in a BlockStatement are not in an expression context
		}
	case *parser.LoopStatement:
		c.compileLoopStatement(node)
	case *parser.RepeatStatement:
		c.compileRepeatStatement(node)
	case *parser.BreakStatement:
		c.compileBreakStatement()
	case *parser.ContinueStatement:
		c.compileContinueStatement()
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
		c.emit(vm.OpPush, vm.VMDataObject{Type: vm.INTGER, IntData: node.Value})
	case *parser.RealLiteral:
		c.emit(vm.OpPush, vm.VMDataObject{Type: vm.REAL, FloatData: node.Value})
	case *parser.StringLiteral:
		c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: node.Value})
	case *parser.Boolean:
		c.emit(vm.OpPush, vm.VMDataObject{Type: vm.BOOLEAN, BoolData: node.Value})
	case *parser.NilLiteral:
		c.emit(vm.OpPush, vm.VMDataObject{}) // Nil
	case *parser.Identifier:
		if node.Value == "" {
			panic("compiling empty identifier")
		}
		c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: node.Value})
		c.emit(vm.OpLoadGlobal)
	case *parser.PackLiteral:
		c.compilePackLiteral(node)
	default:
		panic(fmt.Sprintf("Unsupported expression type: %T", expr))
	}
}

func (c *Compiler) compileBlockExpression(node *parser.BlockExpression) {
	if len(node.Statements) == 0 {
		c.emit(vm.OpPush, vm.VMDataObject{}) // Nil for empty block
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
	c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: node.Name.Value})
	c.emit(vm.OpStoreGlobal)
}

func (c *Compiler) compileDimStatement(node *parser.DimStatement) {
	if node.Value == nil {
		c.emit(vm.OpPush, vm.VMDataObject{}) // Push nil
	} else {
		c.compileExpr(node.Value)
	}
	c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: node.Name.Value})
	c.emit(vm.OpStoreGlobal)
}

func (c *Compiler) compileInfixExpr(node *parser.InfixExpression) {
	c.compileExpr(node.Left)
	c.compileExpr(node.Right)

	switch node.Operator {
	case "+":
		c.emit(vm.OpAdd)
	case "-":
		c.emit(vm.OpSub)
	case "*":
		c.emit(vm.OpMul)
	case "/":
		c.emit(vm.OpDiv)
	case "%":
		c.emit(vm.OpMod)
	case "==":
		c.emit(vm.OpCmpEq)
	case "!=":
		c.emit(vm.OpCmpNeq)
	case ">":
		c.emit(vm.OpCmpGt)
	case "<":
		c.emit(vm.OpCmpLt)
	case ">=":
		c.emit(vm.OpCmpGte)
	case "<=":
		c.emit(vm.OpCmpLte)
	case "and":
		c.emit(vm.OpAnd)
	case "or":
		c.emit(vm.OpOr)
	default:
		panic(fmt.Sprintf("Unknown infix operator: %s", node.Operator))
	}
}

func (c *Compiler) compilePrefixExpr(node *parser.PrefixExpression) {
	c.compileExpr(node.Right)
	switch node.Operator {
	case "not":
		c.emit(vm.OpNot)
	case "-":
		c.emit(vm.OpPush, vm.VMDataObject{Type: vm.INTGER, IntData: 0})
		c.emit(vm.OpSub)
	default:
		panic(fmt.Sprintf("Unknown prefix operator: %s", node.Operator))
	}
}

func (c *Compiler) compileIfExpression(node *parser.IfExpression) {
	c.compileExpr(node.Condition)
	jmpIfFalsePos := c.emitWithPlaceholder(vm.OpJmpIfFalse)

	// Consequence
	if len(node.Consequence.Statements) == 0 {
		c.emit(vm.OpPush, vm.VMDataObject{}) // Nil for empty block
	} else {
		// The value of the block is the value of the last statement
		c.compileStmt(node.Consequence.Statements[len(node.Consequence.Statements)-1], true)
	}

	jmpPos := c.emitWithPlaceholder(vm.OpJmp)
	c.patchJump(jmpIfFalsePos)

	if node.Alternative == nil {
		c.emit(vm.OpPush, vm.VMDataObject{}) // Push nil for false case
	} else {
		c.compileExpr(node.Alternative)
	}
	c.patchJump(jmpPos)
}

func (c *Compiler) compileFunctionLiteral(node *parser.FunctionLiteral) {
	c.emit(vm.OpDefFunc, vm.VMDataObject{Type: vm.STRING, StringData: node.Name.Value})

	// TODO: Handle parameters
	for _, pname := range node.Parameters {
		c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: pname.Value})
		c.emit(vm.OpStoreGlobal)
	}

	c.compileStmt(node.Body, true)
	if len(node.Body.Statements) == 0 || !isReturnStatement(node.Body.Statements[len(node.Body.Statements)-1]) {
		c.emit(vm.OpPush, vm.VMDataObject{}) // Push nil
		c.emit(vm.OpReturn)
	}

}

func (c *Compiler) compileCallExpression(node *parser.CallExpression) {
	for _, arg := range node.Arguments {
		c.compileExpr(arg)
	}

	ident, ok := node.Function.(*parser.Identifier)
	if !ok {
		panic("Calling non-identifier function is not supported yet")
	}

	c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: ident.Value})
	c.emit(vm.OpCall)
}

func (c *Compiler) compileAssignmentExpression(node *parser.AssignmentExpression) {
	if ident, ok := node.Left.(*parser.Identifier); ok {
		c.compileExpr(node.Value)
		c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: ident.Value})
		c.emit(vm.OpStoreGlobal)
	} else if indexExpr, ok := node.Left.(*parser.IndexExpression); ok {
		c.compileExpr(indexExpr.Left)  // pack
		c.compileExpr(indexExpr.Index) // index
		c.compileExpr(node.Value)      // value
		c.emit(vm.OpSetIndex)          // returns modified pack

		if ident, ok := indexExpr.Left.(*parser.Identifier); ok {
			// The modified pack is on the stack. Now push the name and store.
			c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: ident.Value})
			c.emit(vm.OpStoreGlobal)
		} else {
			// The pack was a result of an expression, can't store it back.
			// Pop the modified pack from the stack as it's not used.
			c.emit(vm.OpPop)
		}
	} else if memberAccessExpr, ok := node.Left.(*parser.MemberAccessExpression); ok {
		c.compileExpr(memberAccessExpr.Object)                                                         // pack
		c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: memberAccessExpr.Member.Value}) // index
		c.compileExpr(node.Value)                                                                      // value
		c.emit(vm.OpSetIndex)                                                                          // returns modified pack

		if ident, ok := memberAccessExpr.Object.(*parser.Identifier); ok {
			// The modified pack is on the stack. Now push the name and store.
			c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: ident.Value})
			c.emit(vm.OpStoreGlobal)
		} else {
			c.emit(vm.OpPop)
		}
	} else {
		panic(fmt.Sprintf("Assignment to this expression type is not supported: %T", node.Left))
	}
}

func (c *Compiler) compileLoopStatement(node *parser.LoopStatement) {
	loopStart := len(c.instructions)
	c.pushLoopContext(loopStart)

	c.compileExpr(node.Condition)
	jmpIfFalsePos := c.emitWithPlaceholder(vm.OpJmpIfFalse)

	c.compileStmt(node.Body, false)

	c.emit(vm.OpJmp, vm.VMDataObject{Type: vm.INTGER, IntData: int64(loopStart)})

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
	patchPos := c.emitWithPlaceholder(vm.OpJmp)
	currentLoop := c.loopContexts[len(c.loopContexts)-1]
	currentLoop.breakPatches = append(currentLoop.breakPatches, patchPos)
}

func (c *Compiler) compileContinueStatement() {
	if len(c.loopContexts) == 0 {
		panic("'continue' outside of a loop")
	}
	currentLoop := c.loopContexts[len(c.loopContexts)-1]
	c.emit(vm.OpJmp, vm.VMDataObject{Type: vm.INTGER, IntData: int64(currentLoop.startPos)})
}

func (c *Compiler) compilePackLiteral(node *parser.PackLiteral) {
	//c.emit(vm.OpPush, vm.VMDataObject{Type: vm.INTGER, IntData: int64(len(node.Pairs))})
	c.emit(vm.OpMakePack)

	for key, value := range node.Pairs {
		c.compileExpr(key)
		c.compileExpr(value)
		c.emit(vm.OpSetIndex)
	}
}

func (c *Compiler) compileIndexExpression(node *parser.IndexExpression) {
	c.compileExpr(node.Left)
	c.compileExpr(node.Index)
	c.emit(vm.OpIndex)
}

func (c *Compiler) compileMemberAccessExpression(node *parser.MemberAccessExpression) {
	c.compileExpr(node.Object)
	c.emit(vm.OpPush, vm.VMDataObject{Type: vm.STRING, StringData: node.Member.Value})
	c.emit(vm.OpIndex)
}

// --- Helper methods ---

func (c *Compiler) emit(op vm.VMOp, operands ...vm.VMDataObject) {
	instr := vm.VMInstr{Op: op}
	if len(operands) > 0 {
		instr.Oprand1 = operands[0]
	}
	c.instructions = append(c.instructions, instr)
}

func (c *Compiler) emitWithPlaceholder(op vm.VMOp) int {
	instr := vm.VMInstr{Op: op, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: -1}}
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

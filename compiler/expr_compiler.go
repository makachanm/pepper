package compiler

import (
	"fmt"
	"pepper/parser"
	"pepper/runtime"
)

func (c *Compiler) compileExpr(expr parser.Expression) {
	if expr == nil {
		fmt.Println("Warning: compileExpr called with nil expression, ", expr)
		return
	}
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
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.INTGER, Value: node.Value})
	case *parser.RealLiteral:
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.REAL, Value: node.Value})
	case *parser.StringLiteral:
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, Value: node.Value})
	case *parser.Boolean:
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.BOOLEAN, Value: node.Value})
	case *parser.NilLiteral:
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.NIL, Value: nil}) // Nil
	case *parser.Identifier:
		symbol, ok := c.symbolTable.Resolve(node.Value)
		if !ok {
			token := node.GetToken()
			panic(fmt.Sprintf("line %d:%d: undefined identifier: %s", token.Line, token.Column, node.Value))
		}

		switch symbol.Type {
		case FuncSymbol:
			c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.FUNCTION_ALIAS, Value: symbol.Name})
		case VarSymbol:
			c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, Value: symbol.Name})
			if symbol.Scope == GlobalScope {
				c.emit(runtime.OpLoadGlobal)
			} else {
				c.emit(runtime.OpLoadLocal)
			}
		}
	case *parser.PackLiteral:
		c.compilePackLiteral(node)
	default:
		token := expr.GetToken()
		panic(fmt.Sprintf("line %d:%d: Unsupported expression type: %T", token.Line, token.Column, expr))
	}
}

func (c *Compiler) compileBlockExpression(node *parser.BlockExpression) {
	if len(node.Statements) == 0 {
		return
	}
	last := len(node.Statements) - 1
	for i, s := range node.Statements {
		isLast := (i == last)
		c.compileStmt(s, isLast) // Only the last statement is in an expression context
	}
}

func (c *Compiler) compileInfixExpr(node *parser.InfixExpression) {
	// --- Constant Folding for Integers and Reals ---
	var leftVal, rightVal float64
	var leftIsNumber, rightIsNumber bool
	var leftIsInt, rightIsInt bool

	if l, ok := node.Left.(*parser.IntegerLiteral); ok {
		leftVal = float64(l.Value)
		leftIsNumber = true
		leftIsInt = true
	} else if l, ok := node.Left.(*parser.RealLiteral); ok {
		leftVal = l.Value
		leftIsNumber = true
	}

	if r, ok := node.Right.(*parser.IntegerLiteral); ok {
		rightVal = float64(r.Value)
		rightIsNumber = true
		rightIsInt = true
	} else if r, ok := node.Right.(*parser.RealLiteral); ok {
		rightVal = r.Value
		rightIsNumber = true
	}

	if leftIsNumber && rightIsNumber {
		var result float64
		isResultInt := leftIsInt && rightIsInt

		switch node.Operator {
		case "+":
			result = leftVal + rightVal
		case "-":
			result = leftVal - rightVal
		case "*":
			result = leftVal * rightVal
		case "/":
			if rightVal == 0.0 {
				goto fallback // Division by zero, cannot fold.
			}
			result = leftVal / rightVal
			isResultInt = false // Division of literals is treated as real division.
		default:
			goto fallback // Operator not supported for folding.
		}

		if isResultInt {
			c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.INTGER, Value: int64(result)})
		} else {
			c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.REAL, Value: result})
		}
		return
	}

fallback:
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
		token := node.GetToken()
		panic(fmt.Sprintf("line %d:%d: Unknown infix operator: %s", token.Line, token.Column, node.Operator))
	}
}

func (c *Compiler) compilePrefixExpr(node *parser.PrefixExpression) {
	c.compileExpr(node.Right)
	switch node.Operator {
	case "not":
		c.emit(runtime.OpNot)
	case "-":
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.INTGER, Value: int64(0)})
		c.emit(runtime.OpSub)
	default:
		token := node.GetToken()
		panic(fmt.Sprintf("line %d:%d: Unknown prefix operator: %s", token.Line, token.Column, node.Operator))
	}
}

func (c *Compiler) compileIfExpression(node *parser.IfExpression) {
	// Attempt to compile for combined compare-and-jump optimization
	if infix, ok := node.Condition.(*parser.InfixExpression); ok {
		jumpOp := c.getJumpOpForInfix(infix.Operator)
		if jumpOp != 0 {
			c.compileExpr(infix.Left)
			c.compileExpr(infix.Right)

			jmpIfFalsePos := c.emitWithPlaceholder(jumpOp)

			// Consequence
			if len(node.Consequence.Statements) != 0 {
				c.compileStmt(node.Consequence, true)
			}

			jmpOverElsePos := c.emitWithPlaceholder(runtime.OpJmp)
			c.patchJump(jmpIfFalsePos)

			// Alternative
			if node.Alternative != nil {
				c.compileStmt(node.Alternative, true)
			}
			c.patchJump(jmpOverElsePos)
			return
		}
	}

	// Fallback for non-optimizable expressions (like 'and', 'or', or single values)
	c.compileExpr(node.Condition)
	jmpIfFalsePos := c.emitWithPlaceholder(runtime.OpJmpIfFalse)

	// Consequence
	if len(node.Consequence.Statements) != 0 {
		c.compileStmt(node.Consequence, true)
	}

	jmpOverElsePos := c.emitWithPlaceholder(runtime.OpJmp)
	c.patchJump(jmpIfFalsePos)

	// Alternative
	if node.Alternative != nil {
		c.compileStmt(node.Alternative, true)
	}
	c.patchJump(jmpOverElsePos)
}

func (c *Compiler) getJumpOpForInfix(op string) runtime.VMOp {
	switch op {
	case "==":
		return runtime.OpJmpIfNeq // Jump if NOT equal
	case "!=":
		return runtime.OpJmpIfEq // Jump if equal
	case ">":
		return runtime.OpJmpIfLte // Jump if NOT greater than (less than or equal)
	case "<":
		return runtime.OpJmpIfGte // Jump if NOT less than (greater than or equal)
	case ">=":
		return runtime.OpJmpIfLt // Jump if less than
	case "<=":
		return runtime.OpJmpIfGt // Jump if greater than
	}
	return 0
}

func (c *Compiler) compileMatchExpression(node *parser.MatchExpression) {
	endJumps := []int{}
	var defaultCase *parser.BlockStatement

	// Find default case
	for _, caseClause := range node.Cases {
		if caseClause.Pattern == nil {
			defaultCase = caseClause.Body
			break
		}
	}

	for _, caseClause := range node.Cases {
		if caseClause.Pattern == nil {
			continue // Skip default case for now
		}

		// Condition
		c.compileExpr(node.Expression)
		c.compileExpr(caseClause.Pattern)
		c.emit(runtime.OpCmpEq)
		jmpIfFalsePos := c.emitWithPlaceholder(runtime.OpJmpIfFalse)

		// Body
		c.compileStmt(caseClause.Body, true)
		endJumps = append(endJumps, c.emitWithPlaceholder(runtime.OpJmp))

		c.patchJump(jmpIfFalsePos)
	}

	// Default case
	if defaultCase != nil {
		c.compileStmt(defaultCase, true)
	}

	// Patch all end jumps
	for _, pos := range endJumps {
		c.patchJump(pos)
	}
}

func (c *Compiler) compileFunctionLiteral(node *parser.FunctionLiteral) {
	c.symbolTable.DefineFunc(node.Name.Value)
	c.emit(runtime.OpDefFunc, runtime.VMDataObject{Type: runtime.STRING, Value: node.Name.Value})
	jumpPos := c.emitWithPlaceholder(runtime.OpJmp)

	c.enterScope()

	for i := len(node.Parameters) - 1; i >= 0; i-- {
		pname := node.Parameters[i]
		c.symbolTable.DefineVar(pname.Value)
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, Value: pname.Value})
		c.emit(runtime.OpStoreLocal)
	}

	c.compileStmt(node.Body, true)
	if len(node.Body.Statements) == 0 || !isReturnStatement(node.Body.Statements[len(node.Body.Statements)-1]) {
		//c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.NIL, Value: nil})
		c.emit(runtime.OpReturn)
	}

	c.leaveScope()
	c.patchJump(jumpPos)
}

func (c *Compiler) compileCallExpression(node *parser.CallExpression) {
	c.usedFunctions[node.Function.(*parser.Identifier).Value] = true
	for _, arg := range node.Arguments {
		c.compileExpr(arg)
	}

	c.compileExpr(node.Function)
	c.emit(runtime.OpCall)
}

func (c *Compiler) compileAssignmentExpression(node *parser.AssignmentExpression) {
	if ident, ok := node.Left.(*parser.Identifier); ok {
		c.compileExpr(node.Value)
		symbol, ok := c.symbolTable.Resolve(ident.Value)
		if !ok {
			token := ident.GetToken()
			panic(fmt.Sprintf("line %d:%d: undefined variable: %s", token.Line, token.Column, ident.Value))
		}
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, Value: symbol.Name})
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
				token := ident.GetToken()
				panic(fmt.Sprintf("line %d:%d: undefined variable: %s", token.Line, token.Column, ident.Value))
			}
			c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, Value: symbol.Name})
			if symbol.Scope == GlobalScope {
				c.emit(runtime.OpStoreGlobal)
			} else {
				c.emit(runtime.OpStoreLocal)
			}
		}
	} else if memberAccessExpr, ok := node.Left.(*parser.MemberAccessExpression); ok {
		c.compileExpr(memberAccessExpr.Object)
		c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, Value: memberAccessExpr.Member.Value})
		c.compileExpr(node.Value)
		c.emit(runtime.OpSetIndex)

		if ident, ok := memberAccessExpr.Object.(*parser.Identifier); ok {
			symbol, ok := c.symbolTable.Resolve(ident.Value)
			if !ok {
				// HACK: Assume global scope if not found
				symbol = Symbol{Name: ident.Value, Scope: GlobalScope}
			}
			c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, Value: symbol.Name})
			if symbol.Scope == GlobalScope {
				c.emit(runtime.OpStoreGlobal)
			} else {
				c.emit(runtime.OpStoreLocal)
			}
		}
	} else {
		token := node.GetToken()
		panic(fmt.Sprintf("line %d:%d: Assignment to this expression type is not supported: %T", token.Line, token.Column, node.Left))
	}
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
	c.emit(runtime.OpPush, runtime.VMDataObject{Type: runtime.STRING, Value: node.Member.Value})
	c.emit(runtime.OpIndex)
}

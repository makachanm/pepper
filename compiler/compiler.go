package compiler

import (
	"pepper/parser"
	"pepper/runtime"
	"pepper/stdfunc"
)

type LoopContext struct {
	startPos     int
	breakPatches []int
}

type Compiler struct {
	instructions         []runtime.VMInstr
	loopContexts         []*LoopContext
	standardFunctionMaps map[string][]runtime.VMInstr
	usedFunctions        map[string]bool
	symbolTable          *SymbolTable
	includedPrograms     map[string]*parser.Program
}

func NewCompiler() *Compiler {
	symbolTable := NewSymbolTable()
	return &Compiler{
		instructions:         make([]runtime.VMInstr, 0, 1024),
		loopContexts:         make([]*LoopContext, 0),
		standardFunctionMaps: make(map[string][]runtime.VMInstr),
		usedFunctions:        make(map[string]bool),
		symbolTable:          symbolTable,
		includedPrograms:     make(map[string]*parser.Program),
	}
}

func (c *Compiler) enterScope() {
	c.symbolTable = NewEnclosedSymbolTable(c.symbolTable)
}

func (c *Compiler) leaveScope() {
	c.symbolTable = c.symbolTable.Outer
}

func (c *Compiler) Compile(program *parser.Program) []runtime.VMInstr {
	c.standardFunctionMaps = stdfunc.DefineStandardFunctions()
	for stdfuncName := range c.standardFunctionMaps {
		c.symbolTable.DefineFunc(stdfuncName)
	}

	for _, stmt := range program.Statements {
		c.compileStmt(stmt, false)
	}

	for stdfuncName, stdInstrs := range c.standardFunctionMaps {
		if c.usedFunctions[stdfuncName] {
			c.instructions = append(c.instructions, runtime.VMInstr{Op: runtime.OpDefFunc, Oprand1: runtime.VMDataObject{Type: runtime.STRING, Value: stdfuncName}})
			c.instructions = append(c.instructions, runtime.VMInstr{Op: runtime.OpJmp, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(len(c.instructions)) + 3}})
			c.instructions = append(c.instructions, stdInstrs...)
			c.instructions = append(c.instructions, runtime.VMInstr{Op: runtime.OpReturn})
		}
	}

	c.optimize()
	return c.instructions
}

func (c *Compiler) optimize() {
	newInstructions := make([]runtime.VMInstr, 0, len(c.instructions))
	i := 0
	for i < len(c.instructions) {
		if i+1 < len(c.instructions) &&
			c.instructions[i].Op == runtime.OpPush &&
			c.instructions[i+1].Op == runtime.OpPop {
			i += 2
			continue
		}
		newInstructions = append(newInstructions, c.instructions[i])
		i++
	}
	c.instructions = newInstructions
}

func (c *Compiler) emit(op runtime.VMOp, operands ...runtime.VMDataObject) {
	instr := runtime.VMInstr{Op: op}
	if len(operands) > 0 {
		instr.Oprand1 = operands[0]
	}
	c.instructions = append(c.instructions, instr)
}

func (c *Compiler) emitWithPlaceholder(op runtime.VMOp) int {
	instr := runtime.VMInstr{Op: op, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(-1)}}
	c.instructions = append(c.instructions, instr)
	return len(c.instructions) - 1
}

func (c *Compiler) patchJump(pos int) {
	jumpTo := len(c.instructions)
	c.instructions[pos].Oprand1.Value = int64(jumpTo)
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
		c.instructions[patchPos].Oprand1.Value = int64(loopEnd)
	}
}

func (c *Compiler) IsStandatrdFunction(name string) bool {
	_, ok := c.standardFunctionMaps[name]
	return ok
}
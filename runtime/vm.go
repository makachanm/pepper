package runtime

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// vmHandler defines the function signature for opcode handlers.
type vmHandler func(v *VM) bool

type VM struct {
	CallStack    *CallStack
	OperandStack *OperandStack
	Memory       *VMMEMObjectTable
	Program      []VMInstr
	PC           int

	curruntFunctionID int

	// dispatchTable holds the handlers for each opcode.
	dispatchTable []vmHandler
	rand          *rand.Rand

	stringTable []string
	stringMap   map[string]int
}

func NewVM(input []VMInstr, wg *sync.WaitGroup) *VM {
	mem := NewVMMEMObjTable()
	GfxNew(640, 480, wg)
	AudioNew()

	vm := &VM{
		CallStack:    NewCallStack(),
		OperandStack: NewOperandStack(),
		Memory:       &mem,
		Program:      input,
		PC:           0,
		rand:         rand.New(rand.NewSource(time.Now().UnixNano())),
		stringMap:    make(map[string]int),
		stringTable:  make([]string, 0),
	}
	vm.curruntFunctionID = vm.internString("")
	vm.initDispatchTable()

	for i, instr := range input {
		if instr.Op == OpDefFunc {
			funcNameID := vm.internString(instr.Oprand1.Value.(string))
			funcObj := VMFunctionObject{
				JumpPc: i + 2,
			}
			vm.Memory.MakeFunc(funcNameID)
			vm.Memory.SetFunc(funcNameID, funcObj, vm)
		}
	}
	return vm
}

// initDispatchTable initializes the dispatch table with opcode handlers.
func (v *VM) initDispatchTable() {
	v.dispatchTable = make([]vmHandler, 256) // Assuming max 256 opcodes
	v.dispatchTable[OpPush] = handlePush
	v.dispatchTable[OpPop] = handlePop
	v.dispatchTable[OpStoreGlobal] = handleStoreGlobal
	v.dispatchTable[OpLoadGlobal] = handleLoadGlobal
	v.dispatchTable[OpStoreLocal] = handleStoreLocal
	v.dispatchTable[OpLoadLocal] = handleLoadLocal
	v.dispatchTable[OpDefFunc] = func(v *VM) bool {
		return true
	}
	v.dispatchTable[OpCall] = handleCall
	v.dispatchTable[OpReturn] = handleReturn
	v.dispatchTable[OpSyscall] = handleSyscall
	v.dispatchTable[OpAdd] = handleAdd
	v.dispatchTable[OpSub] = handleSub
	v.dispatchTable[OpMul] = handleMul
	v.dispatchTable[OpDiv] = handleDiv
	v.dispatchTable[OpMod] = handleMod
	v.dispatchTable[OpAnd] = handleAnd
	v.dispatchTable[OpOr] = handleOr
	v.dispatchTable[OpNot] = handleNot
	v.dispatchTable[OpCmpEq] = handleCmpEq
	v.dispatchTable[OpCmpNeq] = handleCmpNeq
	v.dispatchTable[OpCmpGt] = handleCmpGt
	v.dispatchTable[OpCmpLt] = handleCmpLt
	v.dispatchTable[OpCmpGte] = handleCmpGte
	v.dispatchTable[OpCmpLte] = handleCmpLte
	v.dispatchTable[OpJmp] = handleJmp
	v.dispatchTable[OpJmpIfFalse] = handleJmpIfFalse
	v.dispatchTable[OpJmpIfEq] = handleJmpIfEq
	v.dispatchTable[OpJmpIfNeq] = handleJmpIfNeq
	v.dispatchTable[OpJmpIfGt] = handleJmpIfGt
	v.dispatchTable[OpJmpIfLt] = handleJmpIfLt
	v.dispatchTable[OpJmpIfGte] = handleJmpIfGte
	v.dispatchTable[OpJmpIfLte] = handleJmpIfLte
	v.dispatchTable[OpCstInt] = handleCstInt
	v.dispatchTable[OpCstReal] = handleCstReal
	v.dispatchTable[OpCstStr] = handleCstStr
	v.dispatchTable[OpHlt] = handleHlt
	v.dispatchTable[OpIndex] = handleIndex
	v.dispatchTable[OpMakePack] = handleMakePack
	v.dispatchTable[OpSetIndex] = handleSetIndex
}

func (v *VM) internString(s string) int {
	if id, ok := v.stringMap[s]; ok {
		return id
	}
	id := len(v.stringTable)
	v.stringTable = append(v.stringTable, s)
	v.stringMap[s] = id
	return id
}

func (v *VM) Run(debugmode bool, verboseDebug bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic occurred at PC: %d\n", v.PC)
			DumpMemory(v)
			panic(r) // re-throw panic
		}
	}()

	for v.PC < len(v.Program) {
		if ShouldQuit {
			return
		}
		instr := v.Program[v.PC]
		if verboseDebug {
			fmt.Printf("[%04d] %s\n", v.PC, ResolveVMInstruction(instr))
		}
		handler := v.dispatchTable[instr.Op]
		if handler(v) {
			v.PC++
		}
	}
}

// --- Opcode Handlers ---

func handlePush(v *VM) bool {
	instr := v.Program[v.PC]
	v.OperandStack.Push(instr.Oprand1)
	return true
}

func handlePop(v *VM) bool {
	v.OperandStack.Pop()
	return true
}

func handleStoreGlobal(v *VM) bool {
	nameID := v.internString(v.OperandStack.Pop().Value.(string))
	val := v.OperandStack.Pop()
	globalScopeID := v.internString("")
	if !v.Memory.HasObj(nameID, globalScopeID, v) {
		v.Memory.MakeObj(nameID, globalScopeID)
	}
	v.Memory.SetObj(nameID, val, globalScopeID, v)
	return true
}

func handleLoadGlobal(v *VM) bool {
	nameID := v.internString(v.OperandStack.Pop().Value.(string))
	globalScopeID := v.internString("")
	val := v.Memory.GetObj(nameID, globalScopeID, v)
	v.OperandStack.Push(*val)
	return true
}

func handleStoreLocal(v *VM) bool {
	nameID := v.internString(v.OperandStack.Pop().Value.(string))
	val := v.OperandStack.Pop()
	if !v.Memory.HasObj(nameID, v.curruntFunctionID, v) {
		v.Memory.MakeObj(nameID, v.curruntFunctionID)
	}
	v.Memory.SetObj(nameID, val, v.curruntFunctionID, v)
	return true
}

func handleLoadLocal(v *VM) bool {
	nameID := v.internString(v.OperandStack.Pop().Value.(string))
	val := v.Memory.GetObj(nameID, v.curruntFunctionID, v)
	v.OperandStack.Push(*val)
	return true
}

func handleCall(v *VM) bool {
	callee := v.OperandStack.Pop()
	var funcNameID int

	switch callee.Type {
	case STRING: // Direct call by name
		funcNameID = v.internString(callee.Value.(string))
	case FUNCTION_ALIAS:
		funcNameID = v.internString(callee.Value.(string))
	default:
		panic("Cannot call a non-function")
	}

	function := v.Memory.GetFunc(funcNameID, v)

	v.CallStack.Push(CallStackObject{PC: v.PC, NameID: v.curruntFunctionID})
	v.curruntFunctionID = funcNameID
	v.PC = function.JumpPc // Jump, no PC increment
	return false
}

func handleReturn(v *VM) bool {
	if v.CallStack.IsEmpty() {
		v.PC = len(v.Program) // Halt execution
		return false
	}
	calldata := v.CallStack.Pop()

	// Do not purge memory on recursive calls for now.
	// This is a temporary fix to prevent crashes in functions like fibonacci.
	if v.curruntFunctionID != calldata.NameID {
		PurgeVMMEM(v.Memory, v)
	}

	v.PC = calldata.PC + 1 // Return to instruction after call
	v.curruntFunctionID = calldata.NameID
	return false
}

func handleSyscall(v *VM) bool {
	instr := v.Program[v.PC]
	doSyscall(v, instr.Oprand1.Value.(int64))
	return true
}

func handleAdd(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a + b }, func(a, b int64) int64 { return a + b }, func(a, b string) string { return a + b }))
	return true
}

func handleSub(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a - b }, func(a, b int64) int64 { return a - b }, nil))
	return true
}

func handleMul(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a * b }, func(a, b int64) int64 { return a * b }, nil))
	return true
}

func handleDiv(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a / b }, func(a, b int64) int64 { return a / b }, nil))
	return true
}

func handleMod(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Operate(right, nil, func(a, b int64) int64 { return a % b }, nil))
	return true
}

func handleAnd(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Type == BOOLEAN && right.Type == BOOLEAN {
		v.OperandStack.Push(VMDataObject{Type: BOOLEAN, Value: left.Value.(bool) && right.Value.(bool)})
	}
	return true
}

func handleOr(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Type == BOOLEAN && right.Type == BOOLEAN {
		v.OperandStack.Push(VMDataObject{Type: BOOLEAN, Value: left.Value.(bool) || right.Value.(bool)})
	}
	return true
}

func handleNot(v *VM) bool {
	val := v.OperandStack.Pop()
	if val.Type == BOOLEAN {
		v.OperandStack.Push(VMDataObject{Type: BOOLEAN, Value: !val.Value.(bool)})
	}
	return true
}

func handleCmpEq(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(VMDataObject{Type: BOOLEAN, Value: left.IsEqualTo(right)})
	return true
}

func handleCmpNeq(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(VMDataObject{Type: BOOLEAN, Value: left.IsNotEqualTo(right)})
	return true
}

func handleCmpGt(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a > b }, func(a, b int64) bool { return a > b }))
	return true
}

func handleCmpLt(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a < b }, func(a, b int64) bool { return a < b }))
	return true
}

func handleCmpGte(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a >= b }, func(a, b int64) bool { return a >= b }))
	return true
}

func handleCmpLte(v *VM) bool {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a <= b }, func(a, b int64) bool { return a <= b }))
	return true
}

func handleJmp(v *VM) bool {
	instr := v.Program[v.PC]
	v.PC = int(instr.Oprand1.Value.(int64))
	return false
}

func handleJmpIfFalse(v *VM) bool {
	instr := v.Program[v.PC]
	condition := v.OperandStack.Pop()
	if condition.Type == BOOLEAN && !condition.Value.(bool) {
		v.PC = int(instr.Oprand1.Value.(int64))
		return false
	}
	return true
}

func handleJmpIfEq(v *VM) bool {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.IsEqualTo(right) {
		v.PC = int(instr.Oprand1.Value.(int64))
		return false
	}
	return true
}

func handleJmpIfNeq(v *VM) bool {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.IsNotEqualTo(right) {
		v.PC = int(instr.Oprand1.Value.(int64))
		return false
	}
	return true
}

func handleJmpIfGt(v *VM) bool {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Compare(right, func(a, b float64) bool { return a > b }, func(a, b int64) bool { return a > b }).Value.(bool) {
		v.PC = int(instr.Oprand1.Value.(int64))
		return false
	}
	return true
}

func handleJmpIfLt(v *VM) bool {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Compare(right, func(a, b float64) bool { return a < b }, func(a, b int64) bool { return a < b }).Value.(bool) {
		v.PC = int(instr.Oprand1.Value.(int64))
		return false
	}
	return true
}

func handleJmpIfGte(v *VM) bool {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Compare(right, func(a, b float64) bool { return a >= b }, func(a, b int64) bool { return a >= b }).Value.(bool) {
		v.PC = int(instr.Oprand1.Value.(int64))
		return false
	}
	return true
}

func handleJmpIfLte(v *VM) bool {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Compare(right, func(a, b float64) bool { return a <= b }, func(a, b int64) bool { return a <= b }).Value.(bool) {
		v.PC = int(instr.Oprand1.Value.(int64))
		return false
	}
	return true
}

func handleCstInt(v *VM) bool {
	val := v.OperandStack.Pop()
	v.OperandStack.Push(val.CastTo(INTGER))
	return true
}

func handleCstReal(v *VM) bool {
	val := v.OperandStack.Pop()
	v.OperandStack.Push(val.CastTo(REAL))
	return true
}

func handleCstStr(v *VM) bool {
	val := v.OperandStack.Pop()
	v.OperandStack.Push(val.CastTo(STRING))
	return true
}

func handleHlt(v *VM) bool {
	ShouldQuit = true
	v.PC = len(v.Program) // Stop the loop
	return false
}

func handleIndex(v *VM) bool {
	index := v.OperandStack.Pop()
	packObj := v.OperandStack.Pop()
	if packObj.Type != PACK || packObj.Value == nil {
		v.OperandStack.Push(makeNilValueObj())
		return true
	}
	packData := packObj.Value.(map[PackKey]VMDataObject)
	key := PackKey{
		Type:  index.Type,
		Value: index.Value,
	}
	if val, ok := packData[key]; ok {
		v.OperandStack.Push(val)
	} else {
		v.OperandStack.Push(makeNilValueObj())
	}
	return true
}

func handleMakePack(v *VM) bool {
	pack := make(map[PackKey]VMDataObject)
	v.OperandStack.Push(VMDataObject{Type: PACK, Value: pack})
	return true
}

func handleSetIndex(v *VM) bool {
	value := v.OperandStack.Pop()
	index := v.OperandStack.Pop()
	packObj := v.OperandStack.Pop()
	if packObj.Type != PACK || packObj.Value == nil {
		return true
	}
	packData := packObj.Value.(map[PackKey]VMDataObject)
	key := PackKey{
		Type:  index.Type,
		Value: index.Value,
	}
	packData[key] = value
	v.OperandStack.Push(packObj)
	return true
}

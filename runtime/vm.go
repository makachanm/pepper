package runtime

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// vmHandler defines the function signature for opcode handlers.
type vmHandler func(v *VM)

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
	GfxNew(640, 480, wg) // Initialize graphics context

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
	v.dispatchTable[OpDefFunc] = handleDefFunc
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
		if handler != nil {
			handler(v)
		} else {
			panic(fmt.Sprintf("Unsupported opcode: %d at PC: %d", instr.Op, v.PC))
		}
	}
}

// --- Opcode Handlers ---

func handlePush(v *VM) {
	instr := v.Program[v.PC]
	v.OperandStack.Push(instr.Oprand1)
	v.PC++
}

func handlePop(v *VM) {
	v.OperandStack.Pop()
	v.PC++
}

func handleStoreGlobal(v *VM) {
	nameID := v.internString(v.OperandStack.Pop().Value.(string))
	val := v.OperandStack.Pop()
	globalScopeID := v.internString("")
	if !v.Memory.HasObj(nameID, globalScopeID) {
		v.Memory.MakeObj(nameID, globalScopeID)
	}
	v.Memory.SetObj(nameID, val, globalScopeID)
	v.PC++
}

func handleLoadGlobal(v *VM) {
	nameID := v.internString(v.OperandStack.Pop().Value.(string))
	globalScopeID := v.internString("")
	val := v.Memory.GetObj(nameID, globalScopeID)
	v.OperandStack.Push(*val)
	v.PC++
}

func handleStoreLocal(v *VM) {
	nameID := v.internString(v.OperandStack.Pop().Value.(string))
	val := v.OperandStack.Pop()
	if !v.Memory.HasObj(nameID, v.curruntFunctionID) {
		v.Memory.MakeObj(nameID, v.curruntFunctionID)
	}
	v.Memory.SetObj(nameID, val, v.curruntFunctionID)
	v.PC++
}

func handleLoadLocal(v *VM) {
	nameID := v.internString(v.OperandStack.Pop().Value.(string))
	val := v.Memory.GetObj(nameID, v.curruntFunctionID)
	v.OperandStack.Push(*val)
	v.PC++
}

func handleDefFunc(v *VM) {
	instr := v.Program[v.PC]
	funcNameID := v.internString(instr.Oprand1.Value.(string))
	funcObj := VMFunctionObject{
		JumpPc: v.PC + 2,
	}
	v.Memory.MakeFunc(funcNameID)
	v.Memory.SetFunc(funcNameID, funcObj)
	v.PC++
}

func handleCall(v *VM) {
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

	function := v.Memory.GetFunc(funcNameID)

	v.CallStack.Push(CallStackObject{PC: v.PC, NameID: v.curruntFunctionID})
	v.curruntFunctionID = funcNameID
	v.PC = function.JumpPc // Jump, no PC increment
}

func handleReturn(v *VM) {
	if v.CallStack.IsEmpty() {
		v.PC = len(v.Program) // Halt execution
		return
	}
	calldata := v.CallStack.Pop()

	PurgeVMMEM(v.Memory, v)

	v.PC = calldata.PC + 1 // Return to instruction after call
	v.curruntFunctionID = calldata.NameID
}

func handleSyscall(v *VM) {
	instr := v.Program[v.PC]
	doSyscall(*v, instr.Oprand1.Value.(int64))
	v.PC++
}

func handleAdd(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a + b }, func(a, b int64) int64 { return a + b }, func(a, b string) string { return a + b }))
	v.PC++
}

func handleSub(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a - b }, func(a, b int64) int64 { return a - b }, nil))
	v.PC++
}

func handleMul(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a * b }, func(a, b int64) int64 { return a * b }, nil))
	v.PC++
}

func handleDiv(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a / b }, func(a, b int64) int64 { return a / b }, nil))
	v.PC++
}

func handleMod(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Operate(right, nil, func(a, b int64) int64 { return a % b }, nil))
	v.PC++
}

func handleAnd(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Type == BOOLEAN && right.Type == BOOLEAN {
		v.OperandStack.Push(VMDataObject{Type: BOOLEAN, Value: left.Value.(bool) && right.Value.(bool)})
	}
	v.PC++
}

func handleOr(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Type == BOOLEAN && right.Type == BOOLEAN {
		v.OperandStack.Push(VMDataObject{Type: BOOLEAN, Value: left.Value.(bool) || right.Value.(bool)})
	}
	v.PC++
}

func handleNot(v *VM) {
	val := v.OperandStack.Pop()
	if val.Type == BOOLEAN {
		v.OperandStack.Push(VMDataObject{Type: BOOLEAN, Value: !val.Value.(bool)})
	}
	v.PC++
}

func handleCmpEq(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(VMDataObject{Type: BOOLEAN, Value: left.IsEqualTo(right)})
	v.PC++
}

func handleCmpNeq(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(VMDataObject{Type: BOOLEAN, Value: left.IsNotEqualTo(right)})
	v.PC++
}

func handleCmpGt(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a > b }, func(a, b int64) bool { return a > b }))
	v.PC++
}

func handleCmpLt(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a < b }, func(a, b int64) bool { return a < b }))
	v.PC++
}

func handleCmpGte(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a >= b }, func(a, b int64) bool { return a >= b }))
	v.PC++
}

func handleCmpLte(v *VM) {
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a <= b }, func(a, b int64) bool { return a <= b }))
	v.PC++
}

func handleJmp(v *VM) {
	instr := v.Program[v.PC]
	v.PC = int(instr.Oprand1.Value.(int64))
}

func handleJmpIfFalse(v *VM) {
	instr := v.Program[v.PC]
	condition := v.OperandStack.Pop()
	if condition.Type == BOOLEAN && !condition.Value.(bool) {
		v.PC = int(instr.Oprand1.Value.(int64))
	} else {
		v.PC++
	}
}

func handleJmpIfEq(v *VM) {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.IsEqualTo(right) {
		v.PC = int(instr.Oprand1.Value.(int64))
	} else {
		v.PC++
	}
}

func handleJmpIfNeq(v *VM) {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.IsNotEqualTo(right) {
		v.PC = int(instr.Oprand1.Value.(int64))
	} else {
		v.PC++
	}
}

func handleJmpIfGt(v *VM) {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Compare(right, func(a, b float64) bool { return a > b }, func(a, b int64) bool { return a > b }).Value.(bool) {
		v.PC = int(instr.Oprand1.Value.(int64))
	} else {
		v.PC++
	}
}

func handleJmpIfLt(v *VM) {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Compare(right, func(a, b float64) bool { return a < b }, func(a, b int64) bool { return a < b }).Value.(bool) {
		v.PC = int(instr.Oprand1.Value.(int64))
	} else {
		v.PC++
	}
}

func handleJmpIfGte(v *VM) {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Compare(right, func(a, b float64) bool { return a >= b }, func(a, b int64) bool { return a >= b }).Value.(bool) {
		v.PC = int(instr.Oprand1.Value.(int64))
	} else {
		v.PC++
	}
}

func handleJmpIfLte(v *VM) {
	instr := v.Program[v.PC]
	right := v.OperandStack.Pop()
	left := v.OperandStack.Pop()
	if left.Compare(right, func(a, b float64) bool { return a <= b }, func(a, b int64) bool { return a <= b }).Value.(bool) {
		v.PC = int(instr.Oprand1.Value.(int64))
	} else {
		v.PC++
	}
}

func handleCstInt(v *VM) {
	val := v.OperandStack.Pop()
	v.OperandStack.Push(val.CastTo(INTGER))
	v.PC++
}

func handleCstReal(v *VM) {
	val := v.OperandStack.Pop()
	v.OperandStack.Push(val.CastTo(REAL))
	v.PC++
}

func handleCstStr(v *VM) {
	val := v.OperandStack.Pop()
	v.OperandStack.Push(val.CastTo(STRING))
	v.PC++
}

func handleHlt(v *VM) {
	ShouldQuit = true
	v.PC = len(v.Program) // Stop the loop
}

func handleIndex(v *VM) {
	index := v.OperandStack.Pop()
	packObj := v.OperandStack.Pop()
	if packObj.Type != PACK || packObj.Value == nil {
		v.OperandStack.Push(makeNilValueObj())
		v.PC++
		return
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
	v.PC++
}

func handleMakePack(v *VM) {
	pack := make(map[PackKey]VMDataObject)
	v.OperandStack.Push(VMDataObject{Type: PACK, Value: pack})
	v.PC++
}

func handleSetIndex(v *VM) {
	value := v.OperandStack.Pop()
	index := v.OperandStack.Pop()
	packObj := v.OperandStack.Pop()
	if packObj.Type != PACK || packObj.Value == nil {
		v.PC++
		return
	}
	packData := packObj.Value.(map[PackKey]VMDataObject)
	key := PackKey{
		Type:  index.Type,
		Value: index.Value,
	}
	packData[key] = value
	v.OperandStack.Push(packObj)
	v.PC++
}

package runtime

import (
	"fmt"
	"sync"
)

type VM struct {
	CallStack    *CallStack
	OperandStack *OperandStack
	Memory       *VMMEMObjectTable
	Program      []VMInstr
	PC           int

	curruntFunctionName string
	callDepth           int
}

func NewVM(input []VMInstr, wg *sync.WaitGroup) *VM {
	mem := NewVMMEMObjTable()
	GfxNew(640, 480, wg) // Initialize graphics context

	return &VM{
		CallStack:    NewCallStack(),
		OperandStack: NewOperandStack(),
		Memory:       &mem,
		Program:      input,
		PC:           0,

		curruntFunctionName: "",
		callDepth:           0,
	}
}

func (v *VM) Run(debugmode bool) {
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

		switch instr.Op {
		case OpPush:
			v.OperandStack.Push(instr.Oprand1)
		case OpPop:
			v.OperandStack.Pop()
		case OpStoreGlobal:
			name := v.OperandStack.Pop().StringData
			val := v.OperandStack.Pop()
			if !v.Memory.HasObj(name, "") {
				v.Memory.MakeObj(name, "")
			}
			v.Memory.SetObj(name, val, "")
		case OpLoadGlobal:
			name := v.OperandStack.Pop().StringData
			val := v.Memory.GetObj(name, "")
			v.OperandStack.Push(*val)
		case OpStoreLocal:
			name := v.OperandStack.Pop().StringData
			val := v.OperandStack.Pop()
			if !v.Memory.HasObj(name, v.curruntFunctionName) {
				v.Memory.MakeObj(name, v.curruntFunctionName)
			}
			v.Memory.SetObj(name, val, v.curruntFunctionName)
		case OpLoadLocal:
			name := v.OperandStack.Pop().StringData
			val := v.Memory.GetObj(name, v.curruntFunctionName)
			v.OperandStack.Push(*val)
		case OpDefFunc:
			funcName := instr.Oprand1.StringData
			funcObj := VMFunctionObject{
				JumpPc: v.PC + 2,
			}
			v.Memory.MakeFunc(funcName)
			v.Memory.SetFunc(funcName, funcObj)
		case OpCall:
			funcName := v.OperandStack.Pop().StringData
			function := v.Memory.GetFunc(funcName)

			v.CallStack.Push(CallStackObject{PC: v.PC, Name: v.curruntFunctionName})
			v.PC = function.JumpPc
			v.callDepth++
			v.curruntFunctionName = funcName
			continue

		case OpReturn:
			if v.CallStack.IsEmpty() {
				return // Or handle error
			}
			calldata := v.CallStack.Pop()

			PurgeVMMEM(v.Memory, v)

			v.PC = calldata.PC
			v.curruntFunctionName = calldata.Name
			v.callDepth--

		case OpSyscall:
			doSyscall(*v, instr.Oprand1.IntData)
		case OpAdd:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a + b }, func(a, b int64) int64 { return a + b }, func(a, b string) string { return a + b }))
		case OpSub:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a - b }, func(a, b int64) int64 { return a - b }, nil))
		case OpMul:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a * b }, func(a, b int64) int64 { return a * b }, nil))
		case OpDiv:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a / b }, func(a, b int64) int64 { return a / b }, nil))
		case OpMod:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Operate(right, nil, func(a, b int64) int64 { return a % b }, nil))
		case OpAnd:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			if left.Type == BOOLEAN && right.Type == BOOLEAN {
				v.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: left.BoolData && right.BoolData})
			}
		case OpOr:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			if left.Type == BOOLEAN && right.Type == BOOLEAN {
				v.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: left.BoolData || right.BoolData})
			}
		case OpNot:
			val := v.OperandStack.Pop()
			if val.Type == BOOLEAN {
				v.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: !val.BoolData})
			}
		case OpCmpEq:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: left.IsEqualTo(right)})
		case OpCmpNeq:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: left.IsNotEqualTo(right)})
		case OpCmpGt:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a > b }, func(a, b int64) bool { return a > b }))
		case OpCmpLt:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a < b }, func(a, b int64) bool { return a < b }))
		case OpCmpGte:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a >= b }, func(a, b int64) bool { return a >= b }))
		case OpCmpLte:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a <= b }, func(a, b int64) bool { return a <= b }))
		case OpJmp:
			v.PC = int(instr.Oprand1.IntData)
			continue
		case OpJmpIfFalse:
			condition := v.OperandStack.Pop()
			if condition.Type == BOOLEAN && !condition.BoolData {
				v.PC = int(instr.Oprand1.IntData)
				continue
			}
		case OpJmpIfEq:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			if left.IsEqualTo(right) {
				v.PC = int(instr.Oprand1.IntData)
				continue
			}
		case OpJmpIfNeq:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			if left.IsNotEqualTo(right) {
				v.PC = int(instr.Oprand1.IntData)
				continue
			}
		case OpJmpIfGt:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			if left.Compare(right, func(a, b float64) bool { return a > b }, func(a, b int64) bool { return a > b }).BoolData {
				v.PC = int(instr.Oprand1.IntData)
				continue
			}
		case OpJmpIfLt:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			if left.Compare(right, func(a, b float64) bool { return a < b }, func(a, b int64) bool { return a < b }).BoolData {
				v.PC = int(instr.Oprand1.IntData)
				continue
			}
		case OpJmpIfGte:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			if left.Compare(right, func(a, b float64) bool { return a >= b }, func(a, b int64) bool { return a >= b }).BoolData {
				v.PC = int(instr.Oprand1.IntData)
				continue
			}
		case OpJmpIfLte:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			if left.Compare(right, func(a, b float64) bool { return a <= b }, func(a, b int64) bool { return a <= b }).BoolData {
				v.PC = int(instr.Oprand1.IntData)
				continue
			}
		case OpCstInt:
			val := v.OperandStack.Pop()
			v.OperandStack.Push(val.CastTo(INTGER))
		case OpCstReal:
			val := v.OperandStack.Pop()
			v.OperandStack.Push(val.CastTo(REAL))
		case OpCstStr:
			val := v.OperandStack.Pop()
			v.OperandStack.Push(val.CastTo(STRING))
		case OpHlt:
			ShouldQuit = true
			return
		case OpIndex:
			index := v.OperandStack.Pop()
			pack := v.OperandStack.Pop()
			if pack.Type != PACK || pack.PackData == nil {
				v.OperandStack.Push(VMDataObject{}) // Push nil
				break
			}
			key := PackKey{
				Type:       index.Type,
				IntData:    index.IntData,
				FloatData:  index.FloatData,
				BoolData:   index.BoolData,
				StringData: index.StringData,
			}
			if val, ok := (pack.PackData)[key]; ok {
				v.OperandStack.Push(val)
			} else {
				v.OperandStack.Push(VMDataObject{}) // Push nil
			}
		case OpMakePack:
			pack := make(map[PackKey]VMDataObject)
			v.OperandStack.Push(VMDataObject{Type: PACK, PackData: pack})
		case OpSetIndex:
			value := v.OperandStack.Pop()
			index := v.OperandStack.Pop()
			pack := v.OperandStack.Pop()
			if pack.Type != PACK || pack.PackData == nil {
				break
			}
			key := PackKey{
				Type:       index.Type,
				IntData:    index.IntData,
				FloatData:  index.FloatData,
				BoolData:   index.BoolData,
				StringData: index.StringData,
			}
			(pack.PackData)[key] = value
			v.OperandStack.Push(pack)

		}

		v.PC++
	}
}


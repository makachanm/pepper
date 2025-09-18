package runtime

import (
	"pepper/vm"
)

type VM struct {
	CallStack    *vm.CallStack
	OperandStack *vm.OperandStack
	Memory       *vm.VMMEMObjectTable
	Program      []vm.VMInstr
	PC           int
}

func NewVM(input []vm.VMInstr) *VM {
	mem := vm.NewVMMEMObjTable()
	GfxNew(640, 480) // Initialize graphics context
	return &VM{
		CallStack:    vm.NewCallStack(),
		OperandStack: vm.NewOperandStack(),
		Memory:       &mem,
		Program:      input,
		PC:           0,
	}
}

func (v *VM) Run() {
	for v.PC < len(v.Program) {
		instr := v.Program[v.PC]

		switch instr.Op {
		case vm.OpPush:
			v.OperandStack.Push(instr.Oprand1)
		case vm.OpPop:
			v.OperandStack.Pop()
		case vm.OpStoreGlobal:
			name := v.OperandStack.Pop().StringData
			val := v.OperandStack.Pop()
			if !v.Memory.HasObj(name) {
				v.Memory.MakeObj(name)
			}
			v.Memory.SetObj(name, val)
		case vm.OpLoadGlobal:
			name := v.OperandStack.Pop().StringData
			val := v.Memory.GetObj(name)
			v.OperandStack.Push(*val)
		case vm.OpDefFunc:
			funcName := instr.Oprand1.StringData
			funcObj := vm.VMFunctionObject{
				JumpPc: v.PC,
			}
			v.Memory.MakeFunc(funcName)
			v.Memory.SetFunc(funcName, funcObj)

			// Skip to the end of the function definition
			for v.PC < len(v.Program) && v.Program[v.PC].Op != vm.OpReturn {
				v.PC += 2
			}
		case vm.OpCall:
			funcName := v.OperandStack.Pop().StringData
			function := v.Memory.GetFunc(funcName)

			v.CallStack.Push(v.PC)
			v.PC = function.JumpPc
		case vm.OpReturn:
			if v.CallStack.IsEmpty() {
				return // Or handle error
			}
			v.PC = v.CallStack.Pop()
		case vm.OpSyscall:
			doSyscall(*v, instr.Oprand1.IntData)
		case vm.OpAdd:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a + b }, func(a, b int64) int64 { return a + b }, func(a, b string) string { return a + b }))
		case vm.OpSub:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a - b }, func(a, b int64) int64 { return a - b }, nil))
		case vm.OpMul:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a * b }, func(a, b int64) int64 { return a * b }, nil))
		case vm.OpDiv:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Operate(right, func(a, b float64) float64 { return a / b }, func(a, b int64) int64 { return a / b }, nil))
		case vm.OpMod:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Operate(right, nil, func(a, b int64) int64 { return a % b }, nil))
		case vm.OpAnd:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			if left.Type == vm.BOOLEAN && right.Type == vm.BOOLEAN {
				v.OperandStack.Push(vm.VMDataObject{Type: vm.BOOLEAN, BoolData: left.BoolData && right.BoolData})
			}
		case vm.OpOr:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			if left.Type == vm.BOOLEAN && right.Type == vm.BOOLEAN {
				v.OperandStack.Push(vm.VMDataObject{Type: vm.BOOLEAN, BoolData: left.BoolData || right.BoolData})
			}
		case vm.OpNot:
			val := v.OperandStack.Pop()
			if val.Type == vm.BOOLEAN {
				v.OperandStack.Push(vm.VMDataObject{Type: vm.BOOLEAN, BoolData: !val.BoolData})
			}
		case vm.OpCmpEq:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(vm.VMDataObject{Type: vm.BOOLEAN, BoolData: left.IsEqualTo(right)})
		case vm.OpCmpNeq:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(vm.VMDataObject{Type: vm.BOOLEAN, BoolData: left.IsNotEqualTo(right)})
		case vm.OpCmpGt:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a > b }, func(a, b int64) bool { return a > b }))
		case vm.OpCmpLt:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a < b }, func(a, b int64) bool { return a < b }))
		case vm.OpCmpGte:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a >= b }, func(a, b int64) bool { return a >= b }))
		case vm.OpCmpLte:
			right := v.OperandStack.Pop()
			left := v.OperandStack.Pop()
			v.OperandStack.Push(left.Compare(right, func(a, b float64) bool { return a <= b }, func(a, b int64) bool { return a <= b }))
		case vm.OpJmp:
			v.PC = int(instr.Oprand1.IntData)
			continue
		case vm.OpJmpIfFalse:
			condition := v.OperandStack.Pop()
			if condition.Type == vm.BOOLEAN && !condition.BoolData {
				v.PC = int(instr.Oprand1.IntData)
				continue
			}
		case vm.OpCstInt:
			val := v.OperandStack.Pop()
			v.OperandStack.Push(val.CastTo(vm.INTGER))
		case vm.OpCstReal:
			val := v.OperandStack.Pop()
			v.OperandStack.Push(val.CastTo(vm.REAL))
		case vm.OpCstStr:
			val := v.OperandStack.Pop()
			v.OperandStack.Push(val.CastTo(vm.STRING))
		case vm.OpHlt:
			return
		case vm.OpIndex:
			index := v.OperandStack.Pop()
			pack := v.OperandStack.Pop()
			if pack.Type != vm.PACK || pack.PackData == nil {
				v.OperandStack.Push(vm.VMDataObject{}) // Push nil
				break
			}
			key := vm.PackKey{
				Type:       index.Type,
				IntData:    index.IntData,
				FloatData:  index.FloatData,
				BoolData:   index.BoolData,
				StringData: index.StringData,
			}
			if val, ok := (*pack.PackData)[key]; ok {
				v.OperandStack.Push(val)
			} else {
				v.OperandStack.Push(vm.VMDataObject{}) // Push nil
			}
		case vm.OpMakePack:
			pack := make(map[vm.PackKey]vm.VMDataObject)
			v.OperandStack.Push(vm.VMDataObject{Type: vm.PACK, PackData: &pack})
		case vm.OpSetIndex:
			value := v.OperandStack.Pop()
			index := v.OperandStack.Pop()
			pack := v.OperandStack.Pop()
			if pack.Type != vm.PACK || pack.PackData == nil {
				break
			}
			key := vm.PackKey{
				Type:       index.Type,
				IntData:    index.IntData,
				FloatData:  index.FloatData,
				BoolData:   index.BoolData,
				StringData: index.StringData,
			}
			(*pack.PackData)[key] = value
			v.OperandStack.Push(pack)

		}

		v.PC++
	}
}

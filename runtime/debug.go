package runtime

import (
	"fmt"
	"pepper/vm"
	"strings"
)

func DumpOperandStack(v *VM) {
	fmt.Println(" ----- OPERAND STACK ----- ")
	stack := v.OperandStack.GetStack()
	if len(stack) == 0 {
		fmt.Println("STACK IS EMPTY")
		return
	}
	for i, obj := range stack {
		fmt.Printf("%d: %s\n", i, formatVMDataObject(obj))
	}
}

func DumpMemory(v *VM) {
	fmt.Println(" ----- DATA TABLE ----- ")
	for i, memdata := range v.Memory.DataTable {
		fmt.Println(i, ":", memdata)
	}

	fmt.Println(" ----- FUNCTION TABLE ----- ")
	for i, memdata := range v.Memory.FunctionTable {
		fmt.Println(i, ":", memdata)
	}

	fmt.Println(" ----- DATA MEMORY ----- ")
	for i, memdata := range v.Memory.DataMemory {
		fmt.Println(i, ":", memdata)
	}

	fmt.Println(" ----- FUNCTION MEMORY ----- ")
	for i, memdata := range v.Memory.FunctionMemory {
		fmt.Println(i, ":", memdata)
	}
}

func ResolveVMInstruction(instr vm.VMInstr) string {
	opCode := ""
	switch instr.Op {
	case vm.OpPush:
		opCode = "OpPush"
	case vm.OpPop:
		opCode = "OpPop"
	case vm.OpStoreGlobal:
		opCode = "OpStoreGlobal"
	case vm.OpLoadGlobal:
		opCode = "OpLoadGlobal"
	case vm.OpDefFunc:
		opCode = "OpDefFunc"
	case vm.OpCall:
		opCode = "OpCall"
	case vm.OpReturn:
		opCode = "OpReturn"
	case vm.OpSyscall:
		opCode = "OpSyscall"
	case vm.OpAdd:
		opCode = "OpAdd"
	case vm.OpSub:
		opCode = "OpSub"
	case vm.OpMul:
		opCode = "OpMul"
	case vm.OpDiv:
		opCode = "OpDiv"
	case vm.OpMod:
		opCode = "OpMod"
	case vm.OpAnd:
		opCode = "OpAnd"
	case vm.OpOr:
		opCode = "OpOr"
	case vm.OpNot:
		opCode = "OpNot"
	case vm.OpCmpEq:
		opCode = "OpCmpEq"
	case vm.OpCmpNeq:
		opCode = "OpCmpNeq"
	case vm.OpCmpGt:
		opCode = "OpCmpGt"
	case vm.OpCmpLt:
		opCode = "OpCmpLt"
	case vm.OpCmpGte:
		opCode = "OpCmpGte"
	case vm.OpCmpLte:
		opCode = "OpCmpLte"
	case vm.OpJmp:
		opCode = "OpJmp"
	case vm.OpJmpIfFalse:
		opCode = "OpJmpIfFalse"
	case vm.OpCstInt:
		opCode = "OpCstInt"
	case vm.OpCstReal:
		opCode = "OpCstReal"
	case vm.OpCstStr:
		opCode = "OpCstStr"
	case vm.OpHlt:
		opCode = "OpHlt"
	case vm.OpIndex:
		opCode = "OpIndex"
	case vm.OpMakePack:
		opCode = "OpMakePack"
	case vm.OpSetIndex:
		opCode = "OpSetIndex"
	default:
		opCode = fmt.Sprintf("UnknownOp(%d)", instr.Op)
	}

	// Format operands
	oprand1Str := formatVMDataObject(instr.Oprand1)

	return fmt.Sprintf("%s %s", opCode, oprand1Str)
}

func formatVMDataObject(obj vm.VMDataObject) string {
	switch obj.Type {
	case vm.INTGER:
		return fmt.Sprintf("INT(%d)", obj.IntData)
	case vm.REAL:
		return fmt.Sprintf("REAL(%f)", obj.FloatData)
	case vm.STRING:
		if obj.StringData == "\n" {
			return fmt.Sprintf("STR(%s)", "newline")
		}
		return fmt.Sprintf("STR(%s)", obj.StringData)
	case vm.BOOLEAN:
		return fmt.Sprintf("BOOL(%t)", obj.BoolData)
	case vm.PACK:
		var builder strings.Builder
		builder.WriteString("PACK([")
		i := 0
		for k, v := range obj.PackData {
			builder.WriteString(k.String())
			builder.WriteString(": ")
			builder.WriteString(formatVMDataObject(v))
			if i < len(obj.PackData)-1 {
				builder.WriteString(", ")
			}
			i++
		}
		builder.WriteString("])")
		return builder.String()
	default:
		return "EMPTY"
	}
}

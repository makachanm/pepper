package runtime

import (
	"fmt"
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

func ResolveVMInstruction(instr VMInstr) string {
	opCode := ""
	switch instr.Op {
	case OpPush:
		opCode = "OpPush"
	case OpPop:
		opCode = "OpPop"
	case OpStoreGlobal:
		opCode = "OpStoreGlobal"
	case OpLoadGlobal:
		opCode = "OpLoadGlobal"
	case OpDefFunc:
		opCode = "OpDefFunc"
	case OpCall:
		opCode = "OpCall"
	case OpReturn:
		opCode = "OpReturn"
	case OpSyscall:
		opCode = "OpSyscall"
	case OpAdd:
		opCode = "OpAdd"
	case OpSub:
		opCode = "OpSub"
	case OpMul:
		opCode = "OpMul"
	case OpDiv:
		opCode = "OpDiv"
	case OpMod:
		opCode = "OpMod"
	case OpAnd:
		opCode = "OpAnd"
	case OpOr:
		opCode = "OpOr"
	case OpNot:
		opCode = "OpNot"
	case OpCmpEq:
		opCode = "OpCmpEq"
	case OpCmpNeq:
		opCode = "OpCmpNeq"
	case OpCmpGt:
		opCode = "OpCmpGt"
	case OpCmpLt:
		opCode = "OpCmpLt"
	case OpCmpGte:
		opCode = "OpCmpGte"
	case OpCmpLte:
		opCode = "OpCmpLte"
	case OpJmp:
		opCode = "OpJmp"
	case OpJmpIfFalse:
		opCode = "OpJmpIfFalse"
	case OpCstInt:
		opCode = "OpCstInt"
	case OpCstReal:
		opCode = "OpCstReal"
	case OpCstStr:
		opCode = "OpCstStr"
	case OpHlt:
		opCode = "OpHlt"
	case OpIndex:
		opCode = "OpIndex"
	case OpMakePack:
		opCode = "OpMakePack"
	case OpSetIndex:
		opCode = "OpSetIndex"
	default:
		opCode = fmt.Sprintf("UnknownOp(%d)", instr.Op)
	}

	// Format operands
	oprand1Str := formatVMDataObject(instr.Oprand1)

	return fmt.Sprintf("%s %s", opCode, oprand1Str)
}

func formatVMDataObject(obj VMDataObject) string {
	switch obj.Type {
	case INTGER:
		return fmt.Sprintf("INT(%d)", obj.IntData)
	case REAL:
		return fmt.Sprintf("REAL(%f)", obj.FloatData)
	case STRING:
		if obj.StringData == "\n" {
			return fmt.Sprintf("STR(%s)", "newline")
		}
		return fmt.Sprintf("STR(%s)", obj.StringData)
	case BOOLEAN:
		return fmt.Sprintf("BOOL(%t)", obj.BoolData)
	case PACK:
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

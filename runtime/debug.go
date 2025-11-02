package runtime

import (
	"fmt"
	"strings"
)

func DumpOperandStack(v *VM) {
	fmt.Println("--- Operand Stack ---")
	stack := v.OperandStack.GetStack()
	if len(stack) == 0 {
		fmt.Println("  (empty)")
		return
	}
	for i, obj := range stack {
		fmt.Printf("  %d: %s\n", i, formatVMDataObject(obj))
	}
}

func DumpMemory(v *VM) {
	globalScopeID := v.internString("")

	fmt.Println("--- Globals ---")
	for key, index := range v.Memory.DataTable {
		if key.ScopeKey == globalScopeID {
			name := v.stringTable[key.Name]
			value := v.Memory.DataMemory[index]
			fmt.Printf("  %s = %s\n", name, formatVMDataObject(value))
		}
	}

	if v.curruntFunctionID != globalScopeID {
		currentFunctionName := v.stringTable[v.curruntFunctionID]
		fmt.Printf("--- Locals in %s ---\n", currentFunctionName)
		for key, index := range v.Memory.DataTable {
			if key.ScopeKey == v.curruntFunctionID {
				name := v.stringTable[key.Name]
				value := v.Memory.DataMemory[index]
				fmt.Printf("  %s = %s\n", name, formatVMDataObject(value))
			}
		}
	}

	fmt.Println("--- Functions ---")
	for id, index := range v.Memory.FunctionTable {
		name := v.stringTable[id]
		function := v.Memory.FunctionMemory[index]
		fmt.Printf("  func %s at PC %d\n", name, function.JumpPc)
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
	case OpStoreLocal:
		opCode = "OpStoreLocal"
	case OpLoadLocal:
		opCode = "OpLoadLocal"
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
	case OpJmpIfEq:
		opCode = "OpJmpIfEq"
	case OpJmpIfNeq:
		opCode = "OpJmpIfNeq"
	case OpJmpIfGt:
		opCode = "OpJmpIfGt"
	case OpJmpIfLt:
		opCode = "OpJmpIfLt"
	case OpJmpIfGte:
		opCode = "OpJmpIfGte"
	case OpJmpIfLte:
		opCode = "OpJmpIfLte"
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
	if obj.Value == nil {
		return "EMPTY"
	}
	switch obj.Type {
	case INTGER:
		return fmt.Sprintf("INT(%d)", obj.Value.(int64))
	case REAL:
		return fmt.Sprintf("REAL(%f)", obj.Value.(float64))
	case STRING:
		// Special handling for newline to avoid breaking the line
		strVal := obj.Value.(string)
		if strVal == "\n" {
			return "STR(newline)"
		}
		return fmt.Sprintf("STR(%s)", strVal)
	case BOOLEAN:
		return fmt.Sprintf("BOOL(%t)", obj.Value.(bool))
	case PACK:
		packData, ok := obj.Value.(map[PackKey]VMDataObject)
		if !ok {
			return "PACK(invalid)"
		}
		var builder strings.Builder
		builder.WriteString("PACK([")
		i := 0
		for k, v := range packData {
			builder.WriteString(k.String())
			builder.WriteString(": ")
			builder.WriteString(formatVMDataObject(v))
			if i < len(packData)-1 {
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

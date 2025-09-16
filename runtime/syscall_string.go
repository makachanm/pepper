package runtime

import (
	"pepper/vm"
	"strings"
)

func doSyscallString(v VM, code int64) {
	switch code {
	case 9: // str_len
		str := v.OperandStack.Pop()
		if str.Type == vm.STRING {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.INTGER, IntData: int64(len(str.StringData))})
		} else {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.INTGER, IntData: 0})
		}
	case 10: // str_sub
		end := v.OperandStack.Pop()
		start := v.OperandStack.Pop()
		str := v.OperandStack.Pop()
		if str.Type == vm.STRING && start.Type == vm.INTGER && end.Type == vm.INTGER {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: str.StringData[start.IntData:end.IntData]})
		} else {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
		}
	case 11: // str_replace
		newStr := v.OperandStack.Pop()
		oldStr := v.OperandStack.Pop()
		str := v.OperandStack.Pop()
		if str.Type == vm.STRING && oldStr.Type == vm.STRING && newStr.Type == vm.STRING {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: strings.ReplaceAll(str.StringData, oldStr.StringData, newStr.StringData)})
		} else {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
		}
	}
}

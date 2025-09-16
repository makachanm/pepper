package runtime

import (
	"encoding/json"
	"pepper/vm"
	"strings"
)

func doSyscallString(vmInstance VM, code int64) {
	switch code {
	case 200: // str_len
		str := vmInstance.OperandStack.Pop()
		if str.Type == vm.STRING {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.INTGER, IntData: int64(len(str.StringData))})
		} else {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.INTGER, IntData: 0})
		}
	case 201: // str_sub
		end := vmInstance.OperandStack.Pop()
		start := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == vm.STRING && start.Type == vm.INTGER && end.Type == vm.INTGER {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: str.StringData[start.IntData:end.IntData]})
		} else {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
		}
	case 202: // str_replace
		newStr := vmInstance.OperandStack.Pop()
		oldStr := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == vm.STRING && oldStr.Type == vm.STRING && newStr.Type == vm.STRING {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: strings.ReplaceAll(str.StringData, oldStr.StringData, newStr.StringData)})
		} else {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
		}
	case 203: // json_decode
		jsonStr := vmInstance.OperandStack.Pop()
		if jsonStr.Type != vm.STRING {
			vmInstance.OperandStack.Push(vm.VMDataObject{}) // Push nil
			return
		}
		var data interface{}
		if err := json.Unmarshal([]byte(jsonStr.StringData), &data); err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{}) // Push nil
			return
		}
		vmInstance.OperandStack.Push(convertInterfaceToVMObject(data))
	case 204: // json_encode
		obj := vmInstance.OperandStack.Pop()
		iface := convertVMObjectToInterface(obj)
		jsonBytes, err := json.Marshal(iface)
		if err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: string(jsonBytes)})
	}
}

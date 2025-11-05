package runtime

import (
	"encoding/json"
)

func doSyscallJson(v *VM, code int64) {
	switch code {
	case 500, 204: // json_encode
		val := v.OperandStack.Pop()
		jsonBytes, err := json.Marshal(vmObjectToInterface(val))
		if err != nil {
			v.OperandStack.Push(makeStrValueObj(""))
			return
		}
		v.OperandStack.Push(makeStrValueObj(string(jsonBytes)))
	case 501, 203: // json_decode
		jsonStr := v.OperandStack.Pop()
		if jsonStr.Type != STRING {
			v.OperandStack.Push(makeNilValueObj())
			return
		}
		var result interface{}
		err := json.Unmarshal([]byte(jsonStr.Value.(string)), &result)
		if err != nil {
			v.OperandStack.Push(makeNilValueObj())
			return
		}
		v.OperandStack.Push(interfaceToVmObject(result))
	}
}

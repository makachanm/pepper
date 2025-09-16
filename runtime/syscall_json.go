package runtime

import (
	"encoding/json"
	"pepper/vm"
	"strconv"
)

func doSyscallJson(v VM, code int64) {
	switch code {
	case 500: // json_encode
		val := v.OperandStack.Pop()
		jsonBytes, err := json.Marshal(vmObjectToInterface(val))
		if err != nil {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: string(jsonBytes)})
	case 501: // json_decode
		jsonStr := v.OperandStack.Pop()
		if jsonStr.Type != vm.STRING {
			v.OperandStack.Push(vm.VMDataObject{})
			return
		}
		var result interface{}
		err := json.Unmarshal([]byte(jsonStr.StringData), &result)
		if err != nil {
			v.OperandStack.Push(vm.VMDataObject{})
			return
		}
		v.OperandStack.Push(interfaceToVmObject(result))
	}
}

func vmObjectToInterface(obj vm.VMDataObject) interface{} {
	switch obj.Type {
	case vm.INTGER:
		return obj.IntData
	case vm.REAL:
		return obj.FloatData
	case vm.STRING:
		return obj.StringData
	case vm.BOOLEAN:
		return obj.BoolData
	case vm.PACK:
		if obj.PackData == nil {
			return nil
		}
		// Check if it can be an array
		is_array := true
		max_index := int64(-1)
		for key := range *obj.PackData {
			if key.Type != vm.INTGER {
				is_array = false
				break
			}
			if key.IntData > max_index {
				max_index = key.IntData
			}
		}

		if is_array {
			arr := make([]interface{}, max_index+1)
			for key, val := range *obj.PackData {
				arr[key.IntData] = vmObjectToInterface(val)
			}
			return arr
		} else {
			dict := make(map[string]interface{})
			for key, val := range *obj.PackData {
				dict[key.String()] = vmObjectToInterface(val)
			}
			return dict
		}
	default:
		return nil
	}
}

func interfaceToVmObject(data interface{}) vm.VMDataObject {
	switch v := data.(type) {
	case float64:
		// JSON numbers are float64 by default
		if float64(int64(v)) == v {
			return vm.VMDataObject{Type: vm.INTGER, IntData: int64(v)}
		}
		return vm.VMDataObject{Type: vm.REAL, FloatData: v}
	case string:
		return vm.VMDataObject{Type: vm.STRING, StringData: v}
	case bool:
		return vm.VMDataObject{Type: vm.BOOLEAN, BoolData: v}
	case []interface{}:
		pack := make(map[vm.PackKey]vm.VMDataObject)
		for i, item := range v {
			key := vm.PackKey{Type: vm.INTGER, IntData: int64(i)}
			pack[key] = interfaceToVmObject(item)
		}
		return vm.VMDataObject{Type: vm.PACK, PackData: &pack}
	case map[string]interface{}:
		pack := make(map[vm.PackKey]vm.VMDataObject)
		for k, item := range v {
			key := vm.PackKey{Type: vm.STRING, StringData: k}
			// Attempt to convert key to int if possible, for consistency
			if i, err := strconv.ParseInt(k, 10, 64); err == nil {
				key = vm.PackKey{Type: vm.INTGER, IntData: i}
			}
			pack[key] = interfaceToVmObject(item)
		}
		return vm.VMDataObject{Type: vm.PACK, PackData: &pack}
	default:
		return vm.VMDataObject{}
	}
}

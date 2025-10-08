package runtime

import (
	"encoding/json"
	"strconv"
)

func doSyscallJson(v VM, code int64) {
	switch code {
	case 500: // json_encode
		val := v.OperandStack.Pop()
		jsonBytes, err := json.Marshal(vmObjectToInterface(val))
		if err != nil {
			v.OperandStack.Push(VMDataObject{Type: STRING, StringData: ""})
			return
		}
		v.OperandStack.Push(VMDataObject{Type: STRING, StringData: string(jsonBytes)})
	case 501: // json_decode
		jsonStr := v.OperandStack.Pop()
		if jsonStr.Type != STRING {
			v.OperandStack.Push(VMDataObject{})
			return
		}
		var result interface{}
		err := json.Unmarshal([]byte(jsonStr.StringData), &result)
		if err != nil {
			v.OperandStack.Push(VMDataObject{})
			return
		}
		v.OperandStack.Push(interfaceToVmObject(result))
	}
}

func vmObjectToInterface(obj VMDataObject) interface{} {
	switch obj.Type {
	case INTGER:
		return obj.IntData
	case REAL:
		return obj.FloatData
	case STRING:
		return obj.StringData
	case BOOLEAN:
		return obj.BoolData
	case PACK:
		if obj.PackData == nil {
			return nil
		}
		// Check if it can be an array
		is_array := true
		max_index := int64(-1)
		for key := range obj.PackData {
			if key.Type != INTGER {
				is_array = false
				break
			}
			if key.IntData > max_index {
				max_index = key.IntData
			}
		}

		if is_array {
			arr := make([]interface{}, max_index+1)
			for key, val := range obj.PackData {
				arr[key.IntData] = vmObjectToInterface(val)
			}
			return arr
		} else {
			dict := make(map[string]interface{})
			for key, val := range obj.PackData {
				dict[key.String()] = vmObjectToInterface(val)
			}
			return dict
		}
	default:
		return nil
	}
}

func interfaceToVmObject(data interface{}) VMDataObject {
	switch v := data.(type) {
	case float64:
		// JSON numbers are float64 by default
		if float64(int64(v)) == v {
			return VMDataObject{Type: INTGER, IntData: int64(v)}
		}
		return VMDataObject{Type: REAL, FloatData: v}
	case string:
		return VMDataObject{Type: STRING, StringData: v}
	case bool:
		return VMDataObject{Type: BOOLEAN, BoolData: v}
	case []interface{}:
		pack := make(map[PackKey]VMDataObject)
		for i, item := range v {
			key := PackKey{Type: INTGER, IntData: int64(i)}
			pack[key] = interfaceToVmObject(item)
		}
		return VMDataObject{Type: PACK, PackData: pack}
	case map[string]interface{}:
		pack := make(map[PackKey]VMDataObject)
		for k, item := range v {
			key := PackKey{Type: STRING, StringData: k}
			// Attempt to convert key to int if possible, for consistency
			if i, err := strconv.ParseInt(k, 10, 64); err == nil {
				key = PackKey{Type: INTGER, IntData: i}
			}
			pack[key] = interfaceToVmObject(item)
		}
		return VMDataObject{Type: PACK, PackData: pack}
	default:
		return VMDataObject{}
	}
}

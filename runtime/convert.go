package runtime

import (
	"strconv"
)

func vmObjectToInterface(obj VMDataObject) interface{} {
	switch obj.Type {
	case INTGER, REAL, STRING, BOOLEAN:
		return obj.Value
	case PACK:
		packData, ok := obj.Value.(map[PackKey]VMDataObject)
		if !ok || packData == nil {
			return nil
		}

		// Check if it can be an array
		isArray := true
		if len(packData) > 0 {
			for key := range packData {
				if key.Type != INTGER {
					isArray = false
					break
				}
			}
		} else {
			// Empty pack can be an empty array
			return make([]interface{}, 0)
		}

		if isArray {
			// Check if keys are contiguous from 0
			for i := int64(0); i < int64(len(packData)); i++ {
				if _, ok := packData[PackKey{Type: INTGER, Value: i}]; !ok {
					isArray = false
					break
				}
			}
		}

		if isArray {
			arr := make([]interface{}, len(packData))
			for key, val := range packData {
				arr[key.Value.(int64)] = vmObjectToInterface(val)
			}
			return arr
		} else {
			dict := make(map[string]interface{})
			for key, val := range packData {
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
			return makeIntValueObj(int64(v))
		}
		return makeRealValueObj(v)
	case string:
		return makeStrValueObj(v)
	case bool:
		return makeBoolValueObj(v)
	case []interface{}:
		pack := make(map[PackKey]VMDataObject)
		for i, item := range v {
			key := PackKey{Type: INTGER, Value: int64(i)}
			pack[key] = interfaceToVmObject(item)
		}
		return VMDataObject{Type: PACK, Value: pack}
	case map[string]interface{}:
		pack := make(map[PackKey]VMDataObject)
		for k, item := range v {
			key := PackKey{Type: STRING, Value: k}
			// Attempt to convert key to int if possible, for consistency
			if i, err := strconv.ParseInt(k, 10, 64); err == nil {
				key = PackKey{Type: INTGER, Value: i}
			}
			pack[key] = interfaceToVmObject(item)
		}
		return VMDataObject{Type: PACK, Value: pack}
	default:
		return makeNilValueObj()
	}
}

package runtime

import (
	"sort"
)

// convertInterfaceToVMObject converts a Go interface{} (typically from json.Unmarshal) to a VMDataObject.
func convertInterfaceToVMObject(data interface{}) VMDataObject {
	if data == nil {
		return VMDataObject{} // Represent nil as a nil VMDataObject
	}

	switch v := data.(type) {
	case float64:
		if v == float64(int64(v)) {
			return VMDataObject{Type: INTGER, IntData: int64(v)}
		}
		return VMDataObject{Type: REAL, FloatData: v}
	case string:
		return VMDataObject{Type: STRING, StringData: v}
	case bool:
		return VMDataObject{Type: BOOLEAN, BoolData: v}
	case map[string]interface{}:
		packData := make(map[PackKey]VMDataObject)
		for key, value := range v {
			packKey := PackKey{Type: STRING, StringData: key}
			packData[packKey] = convertInterfaceToVMObject(value)
		}
		return VMDataObject{Type: PACK, PackData: packData}
	case []interface{}:
		packData := make(map[PackKey]VMDataObject)
		for i, value := range v {
			packKey := PackKey{Type: INTGER, IntData: int64(i)}
			packData[packKey] = convertInterfaceToVMObject(value)
		}
		return VMDataObject{Type: PACK, PackData: packData}
	default:
		// Return a nil VMDataObject for unsupported types
		return VMDataObject{}
	}
}

// convertVMObjectToInterface converts a VMDataObject to a Go interface{} for json.Marshal.
func convertVMObjectToInterface(obj VMDataObject) interface{} {
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

		// Check if all keys are integers to determine if it should be an array or map
		isArray := true
		keys := make([]int, 0, len(obj.PackData))
		for k := range obj.PackData {
			if k.Type != INTGER {
				isArray = false
				break
			}
			keys = append(keys, int(k.IntData))
		}

		if isArray {
			// Sort keys to ensure correct order in the slice
			sort.Ints(keys)

			// Check if keys form a sequence starting from 0
			for i, k := range keys {
				if i != k {
					isArray = false
					break
				}
			}

			if isArray {
				slice := make([]interface{}, len(obj.PackData))
				for i := range slice {
					key := PackKey{Type: INTGER, IntData: int64(i)}
					slice[i] = convertVMObjectToInterface((obj.PackData)[key])
				}
				return slice
			}
		}

		// If not a sequential array, treat as a map
		resultMap := make(map[string]interface{})
		for k, v := range obj.PackData {
			resultMap[k.String()] = convertVMObjectToInterface(v)
		}
		return resultMap
	default:
		return nil
	}
}

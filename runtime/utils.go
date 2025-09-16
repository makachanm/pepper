package runtime

import (
	"pepper/vm"
	"sort"
)

// convertInterfaceToVMObject converts a Go interface{} (typically from json.Unmarshal) to a vm.VMDataObject.
func convertInterfaceToVMObject(data interface{}) vm.VMDataObject {
	if data == nil {
		return vm.VMDataObject{} // Represent nil as a nil VMDataObject
	}

	switch v := data.(type) {
	case float64:
		if v == float64(int64(v)) {
			return vm.VMDataObject{Type: vm.INTGER, IntData: int64(v)}
		}
		return vm.VMDataObject{Type: vm.REAL, FloatData: v}
	case string:
		return vm.VMDataObject{Type: vm.STRING, StringData: v}
	case bool:
		return vm.VMDataObject{Type: vm.BOOLEAN, BoolData: v}
	case map[string]interface{}:
		packData := make(map[vm.PackKey]vm.VMDataObject)
		for key, value := range v {
			packKey := vm.PackKey{Type: vm.STRING, StringData: key}
			packData[packKey] = convertInterfaceToVMObject(value)
		}
		return vm.VMDataObject{Type: vm.PACK, PackData: &packData}
	case []interface{}:
		packData := make(map[vm.PackKey]vm.VMDataObject)
		for i, value := range v {
			packKey := vm.PackKey{Type: vm.INTGER, IntData: int64(i)}
			packData[packKey] = convertInterfaceToVMObject(value)
		}
		return vm.VMDataObject{Type: vm.PACK, PackData: &packData}
	default:
		// Return a nil VMDataObject for unsupported types
		return vm.VMDataObject{}
	}
}

// convertVMObjectToInterface converts a vm.VMDataObject to a Go interface{} for json.Marshal.
func convertVMObjectToInterface(obj vm.VMDataObject) interface{} {
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

		// Check if all keys are integers to determine if it should be an array or map
		isArray := true
		keys := make([]int, 0, len(*obj.PackData))
		for k := range *obj.PackData {
			if k.Type != vm.INTGER {
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
				slice := make([]interface{}, len(*obj.PackData))
				for i := range slice {
					key := vm.PackKey{Type: vm.INTGER, IntData: int64(i)}
					slice[i] = convertVMObjectToInterface((*obj.PackData)[key])
				}
				return slice
			}
		}

		// If not a sequential array, treat as a map
		resultMap := make(map[string]interface{})
		for k, v := range *obj.PackData {
			resultMap[k.String()] = convertVMObjectToInterface(v)
		}
		return resultMap
	default:
		return nil
	}
}
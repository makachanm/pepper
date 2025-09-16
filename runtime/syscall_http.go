package runtime

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pepper/vm"
)

// Helper function to convert Go's interface{} from json decoding to a VMDataObject
func convertInterfaceToVMObject(data interface{}) vm.VMDataObject {
	switch v := data.(type) {
	case string:
		return vm.VMDataObject{Type: vm.STRING, StringData: v}
	case float64:
		// JSON numbers are decoded as float64 by default
		return vm.VMDataObject{Type: vm.REAL, FloatData: v}
	case bool:
		return vm.VMDataObject{Type: vm.BOOLEAN, BoolData: v}
	case map[string]interface{}:
		pack := make(map[vm.PackKey]vm.VMDataObject)
		for key, value := range v {
			packKey := vm.PackKey{Type: vm.STRING, StringData: key}
			pack[packKey] = convertInterfaceToVMObject(value)
		}
		return vm.VMDataObject{Type: vm.PACK, PackData: &pack}
	case []interface{}:
		// Represent JSON arrays as Pepper packs with integer keys
		pack := make(map[vm.PackKey]vm.VMDataObject)
		for i, value := range v {
			packKey := vm.PackKey{Type: vm.INTGER, IntData: int64(i)}
			pack[packKey] = convertInterfaceToVMObject(value)
		}
		return vm.VMDataObject{Type: vm.PACK, PackData: &pack}
	case nil:
		return vm.VMDataObject{}
	default:
		// For other types, return a nil object
		return vm.VMDataObject{}
	}
}

func doSyscallHttp(v VM, code int64) {
	switch code {
	case 12: // http_get
		url := v.OperandStack.Pop()
		if url.Type != vm.STRING {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		resp, err := http.Get(url.StringData)
		if err != nil {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: string(body)})
	case 13: // http_post
		body := v.OperandStack.Pop()
		url := v.OperandStack.Pop()
		if url.Type != vm.STRING || body.Type != vm.STRING {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		resp, err := http.Post(url.StringData, "text/plain", bytes.NewBufferString(body.StringData))
		if err != nil {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: string(respBody)})
	case 14: // http_get_json
		url := v.OperandStack.Pop()
		if url.Type != vm.STRING {
			v.OperandStack.Push(vm.VMDataObject{}) // Push nil
			return
		}
		resp, err := http.Get(url.StringData)
		if err != nil {
			v.OperandStack.Push(vm.VMDataObject{}) // Push nil
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			v.OperandStack.Push(vm.VMDataObject{}) // Push nil
			return
		}
		var data interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			v.OperandStack.Push(vm.VMDataObject{}) // Push nil
			return
		}
		v.OperandStack.Push(convertInterfaceToVMObject(data))
	}
}

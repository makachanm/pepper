package runtime

import (
	"encoding/json"
	"io"
	"net/http"
	"pepper/vm"
	"strings"
)

func doSyscallHttp(vmInstance VM, code int64) {
	switch code {
	case 400: // http_get
		url := vmInstance.OperandStack.Pop()
		if url.Type != vm.STRING {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		resp, err := http.Get(url.StringData)
		if err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: string(body)})
	case 401: // http_post
		body := vmInstance.OperandStack.Pop()
		url := vmInstance.OperandStack.Pop()
		if url.Type != vm.STRING || body.Type != vm.STRING {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		resp, err := http.Post(url.StringData, "text/plain", strings.NewReader(body.StringData))
		if err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		defer resp.Body.Close()
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: string(respBody)})
	case 402: // http_get_json
		url := vmInstance.OperandStack.Pop()
		if url.Type != vm.STRING {
			vmInstance.OperandStack.Push(vm.VMDataObject{}) // Push nil
			return
		}
		resp, err := http.Get(url.StringData)
		if err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{}) // Push nil
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{}) // Push nil
			return
		}
		var data interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{}) // Push nil
			return
		}
		vmInstance.OperandStack.Push(convertInterfaceToVMObject(data))
	}
}
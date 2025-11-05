package runtime

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func doSyscallHttp(vmInstance *VM, code int64) {
	switch code {
	case 400: // http_get
		url := vmInstance.OperandStack.Pop()
		if url.Type != STRING {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
			return
		}
		resp, err := http.Get(url.Value.(string))
		if err != nil {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
			return
		}
		vmInstance.OperandStack.Push(makeStrValueObj(string(body)))
	case 401: // http_post
		body := vmInstance.OperandStack.Pop()
		url := vmInstance.OperandStack.Pop()
		if url.Type != STRING || body.Type != STRING {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
			return
		}
		resp, err := http.Post(url.Value.(string), "text/plain", strings.NewReader(body.Value.(string)))
		if err != nil {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
			return
		}
		defer resp.Body.Close()
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
			return
		}
		vmInstance.OperandStack.Push(makeStrValueObj(string(respBody)))
	case 402: // http_get_json
		url := vmInstance.OperandStack.Pop()
		if url.Type != STRING {
			vmInstance.OperandStack.Push(makeNilValueObj())
			return
		}
		resp, err := http.Get(url.Value.(string))
		if err != nil {
			vmInstance.OperandStack.Push(makeNilValueObj())
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			vmInstance.OperandStack.Push(makeNilValueObj())
			return
		}
		var data interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			vmInstance.OperandStack.Push(makeNilValueObj())
			return
		}
		vmInstance.OperandStack.Push(interfaceToVmObject(data))
	}
}

package runtime

import (
	"encoding/json"
	"strings"
)

func doSyscallString(vmInstance VM, code int64) {
	switch code {
	case 200: // str_len
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING {
			vmInstance.OperandStack.Push(VMDataObject{Type: INTGER, IntData: int64(len(str.StringData))})
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: INTGER, IntData: 0})
		}
	case 201: // str_sub
		end := vmInstance.OperandStack.Pop()
		start := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && start.Type == INTGER && end.Type == INTGER {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: str.StringData[start.IntData:end.IntData]})
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: ""})
		}
	case 202: // str_replace
		n := vmInstance.OperandStack.Pop()
		newStr := vmInstance.OperandStack.Pop()
		oldStr := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && oldStr.Type == STRING && newStr.Type == STRING && n.Type == INTGER {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: strings.Replace(str.StringData, oldStr.StringData, newStr.StringData, int(n.IntData))})
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: ""})
		}
	case 203: // json_decode
		jsonStr := vmInstance.OperandStack.Pop()
		if jsonStr.Type != STRING {
			vmInstance.OperandStack.Push(VMDataObject{}) // Push nil
			return
		}
		var data interface{}
		if err := json.Unmarshal([]byte(jsonStr.StringData), &data); err != nil {
			vmInstance.OperandStack.Push(VMDataObject{}) // Push nil
			return
		}
		vmInstance.OperandStack.Push(convertInterfaceToVMObject(data))
	case 204: // json_encode
		obj := vmInstance.OperandStack.Pop()
		iface := convertVMObjectToInterface(obj)
		jsonBytes, err := json.Marshal(iface)
		if err != nil {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: ""})
			return
		}
		vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: string(jsonBytes)})
	case 205: // str_split
		sep := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && sep.Type == STRING {
			parts := strings.Split(str.StringData, sep.StringData)
			elements := VMDataObject{Type: PACK, PackData: map[PackKey]VMDataObject{}}
			for i, part := range parts {
				elements.PackData[PackKey{Type: INTGER, IntData: int64(i)}] = VMDataObject{Type: STRING, StringData: part}
			}
			vmInstance.OperandStack.Push(elements)
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: PACK, PackData: map[PackKey]VMDataObject{}}) // Push empty array
		}
	case 206: // str_join
		sep := vmInstance.OperandStack.Pop()
		arr := vmInstance.OperandStack.Pop()
		if arr.Type == PACK && sep.Type == STRING {
			var parts []string
			for _, obj := range arr.PackData {
				if obj.Type == STRING {
					parts = append(parts, obj.StringData)
				}
			}
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: strings.Join(parts, sep.StringData)})
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: ""})
		}
	case 207: // str_contains
		substr := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && substr.Type == STRING {
			if strings.Contains(str.StringData, substr.StringData) {
				vmInstance.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: true})
			} else {
				vmInstance.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: false})
			}
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: false})
		}
	case 208: // str_has_prefix
		prefix := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && prefix.Type == STRING {
			if strings.HasPrefix(str.StringData, prefix.StringData) {
				vmInstance.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: true})
			} else {
				vmInstance.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: false})
			}
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: false})
		}
	case 209: // str_has_suffix
		suffix := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && suffix.Type == STRING {
			if strings.HasSuffix(str.StringData, suffix.StringData) {
				vmInstance.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: true})
			} else {
				vmInstance.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: false})
			}
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: BOOLEAN, BoolData: false})
		}
	case 210: // str_to_lower
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: strings.ToLower(str.StringData)})
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: ""})
		}
	case 211: // str_to_upper
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: strings.ToUpper(str.StringData)})
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: ""})
		}
	case 212: // str_trim
		cutset := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && cutset.Type == STRING {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: strings.Trim(str.StringData, cutset.StringData)})
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: STRING, StringData: ""})
		}
	case 213: // str_index_of
		substr := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && substr.Type == STRING {
			vmInstance.OperandStack.Push(VMDataObject{Type: INTGER, IntData: int64(strings.Index(str.StringData, substr.StringData))})
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: INTGER, IntData: -1})
		}
	}
}

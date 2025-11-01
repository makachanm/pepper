package runtime

import (
	"strings"
)

func doSyscallString(vmInstance VM, code int64) {
	switch code {
	case 200: // str_len
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING {
			vmInstance.OperandStack.Push(makeIntValueObj(int64(len(str.Value.(string)))))
		} else {
			vmInstance.OperandStack.Push(makeIntValueObj(0))
		}
	case 201: // str_sub
		end := vmInstance.OperandStack.Pop()
		start := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && start.Type == INTGER && end.Type == INTGER {
			vmInstance.OperandStack.Push(makeStrValueObj(str.Value.(string)[start.Value.(int64):end.Value.(int64)]))
		} else {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
		}
	case 202: // str_replace
		n := vmInstance.OperandStack.Pop()
		newStr := vmInstance.OperandStack.Pop()
		oldStr := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && oldStr.Type == STRING && newStr.Type == STRING && n.Type == INTGER {
			vmInstance.OperandStack.Push(makeStrValueObj(strings.Replace(str.Value.(string), oldStr.Value.(string), newStr.Value.(string), int(n.Value.(int64)))))
		} else {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
		}
	case 205: // str_split
		sep := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && sep.Type == STRING {
			parts := strings.Split(str.Value.(string), sep.Value.(string))
			pack := make(map[PackKey]VMDataObject)
			for i, part := range parts {
				pack[PackKey{Type: INTGER, Value: int64(i)}] = makeStrValueObj(part)
			}
			vmInstance.OperandStack.Push(VMDataObject{Type: PACK, Value: pack})
		} else {
			vmInstance.OperandStack.Push(VMDataObject{Type: PACK, Value: make(map[PackKey]VMDataObject)}) // Push empty pack
		}
	case 206: // str_join
		sep := vmInstance.OperandStack.Pop()
		arr := vmInstance.OperandStack.Pop()
		if arr.Type == PACK && sep.Type == STRING {
			var parts []string
			packData := arr.Value.(map[PackKey]VMDataObject)
			// To join in order, we need to sort the keys if they are integers
			// For simplicity, we assume the pack is array-like and iterate up
			for i := 0; i < len(packData); i++ {
				if obj, ok := packData[PackKey{Type: INTGER, Value: int64(i)}]; ok {
					if obj.Type == STRING {
						parts = append(parts, obj.Value.(string))
					}
				}
			}
			vmInstance.OperandStack.Push(makeStrValueObj(strings.Join(parts, sep.Value.(string))))
		} else {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
		}
	case 207: // str_contains
		substr := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && substr.Type == STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(strings.Contains(str.Value.(string), substr.Value.(string))))
		} else {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
		}
	case 208: // str_has_prefix
		prefix := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && prefix.Type == STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(strings.HasPrefix(str.Value.(string), prefix.Value.(string))))
		} else {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
		}
	case 209: // str_has_suffix
		suffix := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && suffix.Type == STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(strings.HasSuffix(str.Value.(string), suffix.Value.(string))))
		} else {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
		}
	case 210: // str_to_lower
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING {
			vmInstance.OperandStack.Push(makeStrValueObj(strings.ToLower(str.Value.(string))))
		} else {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
		}
	case 211: // str_to_upper
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING {
			vmInstance.OperandStack.Push(makeStrValueObj(strings.ToUpper(str.Value.(string))))
		} else {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
		}
	case 212: // str_trim
		cutset := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && cutset.Type == STRING {
			vmInstance.OperandStack.Push(makeStrValueObj(strings.Trim(str.Value.(string), cutset.Value.(string))))
		} else {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
		}
	case 213: // str_index_of
		substr := vmInstance.OperandStack.Pop()
		str := vmInstance.OperandStack.Pop()
		if str.Type == STRING && substr.Type == STRING {
			vmInstance.OperandStack.Push(makeIntValueObj(int64(strings.Index(str.Value.(string), substr.Value.(string)))))
		} else {
			vmInstance.OperandStack.Push(makeIntValueObj(-1))
		}
	}
}
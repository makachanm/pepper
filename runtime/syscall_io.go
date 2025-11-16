package runtime

import (
	"bufio"
	"fmt"
	"os"
)

func doSyscallIO(vmInstance *VM, code int64) {
	switch code {
	case 0: // print
		val := vmInstance.OperandStack.Pop()
		fmt.Print(val.String())
	case 1: // println
		val := vmInstance.OperandStack.Pop()
		fmt.Printf("%s\n", val.String())
	case 2: // io_readln
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		vmInstance.OperandStack.Push(makeStrValueObj(text))
	case 3: // io_read_file
		path := vmInstance.OperandStack.Pop()
		if path.Type != STRING {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
			return
		}
		content, err := os.ReadFile(path.Value.(string))
		if err != nil {
			vmInstance.OperandStack.Push(makeStrValueObj(""))
			return
		}
		vmInstance.OperandStack.Push(makeStrValueObj(string(content)))
	case 4: // io_write_file
		content := vmInstance.OperandStack.Pop()
		path := vmInstance.OperandStack.Pop()
		if path.Type != STRING || content.Type != STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		err := os.WriteFile(path.Value.(string), []byte(content.Value.(string)), 0644)
		if err != nil {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		vmInstance.OperandStack.Push(makeBoolValueObj(true))
	case 5: // file_exists
		path := vmInstance.OperandStack.Pop()
		if path.Type != STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		_, err := os.Stat(path.Value.(string))
		vmInstance.OperandStack.Push(makeBoolValueObj(err == nil))
	case 6: // delete_file
		path := vmInstance.OperandStack.Pop()
		if path.Type != STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		err := os.Remove(path.Value.(string))
		vmInstance.OperandStack.Push(makeBoolValueObj(err == nil))
	case 7: // rename_file
		newPath := vmInstance.OperandStack.Pop()
		oldPath := vmInstance.OperandStack.Pop()
		if oldPath.Type != STRING || newPath.Type != STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}

		err := os.Rename(oldPath.Value.(string), newPath.Value.(string))
		vmInstance.OperandStack.Push(makeBoolValueObj(err == nil))
	case 8: // list_dir
		path := vmInstance.OperandStack.Pop()
		if path.Type != STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		files, err := os.ReadDir(path.Value.(string))
		if err != nil {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		var fileNames []VMDataObject
		for _, file := range files {
			fileNames = append(fileNames, makeStrValueObj(file.Name()))
		}

		pack := make(map[PackKey]VMDataObject)
		for i, name := range fileNames {
			key := PackKey{
				Type:  INTGER,
				Value: int64(i),
			}
			pack[key] = name
		}
		vmInstance.OperandStack.Push(VMDataObject{Type: PACK, Value: pack})
	case 9: // create_dir
		path := vmInstance.OperandStack.Pop()
		if path.Type != STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		err := os.MkdirAll(path.Value.(string), 0755)
		vmInstance.OperandStack.Push(makeBoolValueObj(err == nil))
	case 10: // is_dir
		path := vmInstance.OperandStack.Pop()
		if path.Type != STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		info, err := os.Stat(path.Value.(string))
		if err != nil {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		vmInstance.OperandStack.Push(makeBoolValueObj(info.IsDir()))
	case 11: // is_file
		path := vmInstance.OperandStack.Pop()
		if path.Type != STRING {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		info, err := os.Stat(path.Value.(string))
		if err != nil {
			vmInstance.OperandStack.Push(makeBoolValueObj(false))
			return
		}
		vmInstance.OperandStack.Push(makeBoolValueObj(!info.IsDir()))
	}
}

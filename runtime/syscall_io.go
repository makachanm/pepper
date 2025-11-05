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
	}
}

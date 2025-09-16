package runtime

import (
	"bufio"
	"fmt"
	"os"
	"pepper/vm"
)

func doSyscallIO(vmInstance VM, code int64) {
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
		vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: text})
	case 3: // io_read_file
		path := vmInstance.OperandStack.Pop()
		if path.Type != vm.STRING {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		content, err := os.ReadFile(path.StringData)
		if err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: ""})
			return
		}
		vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.STRING, StringData: string(content)})
	case 4: // io_write_file
		content := vmInstance.OperandStack.Pop()
		path := vmInstance.OperandStack.Pop()
		if path.Type != vm.STRING || content.Type != vm.STRING {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.BOOLEAN, BoolData: false})
			return
		}
		err := os.WriteFile(path.StringData, []byte(content.StringData), 0644)
		if err != nil {
			vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.BOOLEAN, BoolData: false})
			return
		}
		vmInstance.OperandStack.Push(vm.VMDataObject{Type: vm.BOOLEAN, BoolData: true})
	}
}

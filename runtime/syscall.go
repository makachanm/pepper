package runtime

import "fmt"

func doSyscall(vm *VM, code int64) {
	switch code {
	case 0: // Print top of stack
		val := vm.OperandStack.Peek()
		fmt.Print(val.String())
	case 1: // Print and pop top of stack
		val := vm.OperandStack.Pop()
		fmt.Print(val.String())
	}
}

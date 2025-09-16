package runtime

import (
	"fmt"
)

func doSyscallIO(v VM, code int64) {
	switch code {
	case 0: // Print top of stack
		val := v.OperandStack.Pop()
		fmt.Print(val.String())
	case 1: // Print and pop top of stack
		val := v.OperandStack.Pop()
		fmt.Printf("%s\n", val.String())
	}
}

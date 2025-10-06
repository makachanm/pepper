package runtime

import (
	"time"
)

func doSyscallTime(vm VM, code int64) {
	switch code {
	case 600: // sleep
		duration := vm.OperandStack.Pop().IntData
		time.Sleep(time.Duration(duration) * time.Millisecond)
	}
}

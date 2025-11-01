package runtime

import (
	"time"
)

func doSyscallTime(vm VM, code int64) {
	switch code {
	case 600: // sleep
		duration := vm.OperandStack.Pop().Value.(int64)
		time.Sleep(time.Duration(duration) * time.Millisecond)
	case 601: // time.now
		vm.OperandStack.Push(VMDataObject{Type: INTGER, Value: time.Now().Unix()})
	case 602: // time.year
		vm.OperandStack.Push(VMDataObject{Type: INTGER, Value: int64(time.Now().Year())})
	case 603: // time.month
		vm.OperandStack.Push(VMDataObject{Type: INTGER, Value: int64(time.Now().Month())})
	case 604: // time.day
		vm.OperandStack.Push(VMDataObject{Type: INTGER, Value: int64(time.Now().Day())})
	case 605: // time.hour
		vm.OperandStack.Push(VMDataObject{Type: INTGER, Value: int64(time.Now().Hour())})
	case 606: // time.minute
		vm.OperandStack.Push(VMDataObject{Type: INTGER, Value: int64(time.Now().Minute())})
	case 607: // time.second
		vm.OperandStack.Push(VMDataObject{Type: INTGER, Value: int64(time.Now().Second())})
	}
}

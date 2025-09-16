package runtime

import (
	"math"
	"pepper/vm"
)

func doSyscallMath(v VM, code int64) {
	switch code {
	case 100: // sin
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == vm.REAL {
			res = math.Sin(val.FloatData)
		} else if val.Type == vm.INTGER {
			res = math.Sin(float64(val.IntData))
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})
	case 101: // cos
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == vm.REAL {
			res = math.Cos(val.FloatData)
		} else if val.Type == vm.INTGER {
			res = math.Cos(float64(val.IntData))
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})
	case 102: // tan
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == vm.REAL {
			res = math.Tan(val.FloatData)
		} else if val.Type == vm.INTGER {
			res = math.Tan(float64(val.IntData))
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})
	case 103: // sqrt
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == vm.REAL {
			res = math.Sqrt(val.FloatData)
		} else if val.Type == vm.INTGER {
			res = math.Sqrt(float64(val.IntData))
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})
	}
}

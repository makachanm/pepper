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

	case 104: // pow
		exp := v.OperandStack.Pop()
		base := v.OperandStack.Pop()
		var res float64
		var baseVal, expVal float64
		if base.Type == vm.REAL {
			baseVal = base.FloatData
		} else if base.Type == vm.INTGER {
			baseVal = float64(base.IntData)
		}
		if exp.Type == vm.REAL {
			expVal = exp.FloatData
		} else if exp.Type == vm.INTGER {
			expVal = float64(exp.IntData)
		}
		res = math.Pow(baseVal, expVal)
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})

	case 105: // log
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == vm.REAL {
			res = math.Log(val.FloatData)
		} else if val.Type == vm.INTGER {
			res = math.Log(float64(val.IntData))
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})
	case 106: // exp
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == vm.REAL {
			res = math.Exp(val.FloatData)
		} else if val.Type == vm.INTGER {
			res = math.Exp(float64(val.IntData))
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})
	case 107: // abs
		val := v.OperandStack.Pop()
		if val.Type == vm.REAL {
			res := math.Abs(val.FloatData)
			v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})
		} else if val.Type == vm.INTGER {
			res := math.Abs(float64(val.IntData))
			v.OperandStack.Push(vm.VMDataObject{Type: vm.INTGER, IntData: int64(res)})
		}

	case 108: // len of pack
		val := v.OperandStack.Pop()
		if val.Type == vm.PACK {
			length := int64(len(val.PackData))
			v.OperandStack.Push(vm.VMDataObject{Type: vm.INTGER, IntData: length})
		} else {
			panic("len() syscall expects a pack type")
		}

	case 109: // asin
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == vm.REAL {
			res = math.Asin(val.FloatData)
		} else if val.Type == vm.INTGER {
			res = math.Asin(float64(val.IntData))
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})

	case 110: // acos
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == vm.REAL {
			res = math.Acos(val.FloatData)
		} else if val.Type == vm.INTGER {
			res = math.Acos(float64(val.IntData))
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})

	case 111: // atan
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == vm.REAL {
			res = math.Atan(val.FloatData)
		} else if val.Type == vm.INTGER {
			res = math.Atan(float64(val.IntData))
		}
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})

	case 112: // atan2
		x_op := v.OperandStack.Pop()
		y_op := v.OperandStack.Pop()
		var x, y float64
		if x_op.Type == vm.REAL {
			x = x_op.FloatData
		} else if x_op.Type == vm.INTGER {
			x = float64(x_op.IntData)
		}
		if y_op.Type == vm.REAL {
			y = y_op.FloatData
		} else if y_op.Type == vm.INTGER {
			y = float64(y_op.IntData)
		}
		res := math.Atan2(y, x)
		v.OperandStack.Push(vm.VMDataObject{Type: vm.REAL, FloatData: res})
	}
}

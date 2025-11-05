package runtime

import (
	"math"
)

func doSyscallMath(v *VM, code int64) {
	switch code {
	case 100: // sin
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == REAL {
			res = math.Sin(val.Value.(float64))
		} else if val.Type == INTGER {
			res = math.Sin(float64(val.Value.(int64)))
		}
		v.OperandStack.Push(makeRealValueObj(res))
	case 101: // cos
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == REAL {
			res = math.Cos(val.Value.(float64))
		} else if val.Type == INTGER {
			res = math.Cos(float64(val.Value.(int64)))
		}
		v.OperandStack.Push(makeRealValueObj(res))
	case 102: // tan
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == REAL {
			res = math.Tan(val.Value.(float64))
		} else if val.Type == INTGER {
			res = math.Tan(float64(val.Value.(int64)))
		}
		v.OperandStack.Push(makeRealValueObj(res))
	case 103: // sqrt
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == REAL {
			res = math.Sqrt(val.Value.(float64))
		} else if val.Type == INTGER {
			res = math.Sqrt(float64(val.Value.(int64)))
		}
		v.OperandStack.Push(makeRealValueObj(res))

	case 104: // pow
		exp := v.OperandStack.Pop()
		base := v.OperandStack.Pop()
		var res float64
		var baseVal, expVal float64
		if base.Type == REAL {
			baseVal = base.Value.(float64)
		} else if base.Type == INTGER {
			baseVal = float64(base.Value.(int64))
		}
		if exp.Type == REAL {
			expVal = exp.Value.(float64)
		} else if exp.Type == INTGER {
			expVal = float64(exp.Value.(int64))
		}
		res = math.Pow(baseVal, expVal)
		v.OperandStack.Push(makeRealValueObj(res))

	case 105: // log
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == REAL {
			res = math.Log(val.Value.(float64))
		} else if val.Type == INTGER {
			res = math.Log(float64(val.Value.(int64)))
		}
		v.OperandStack.Push(makeRealValueObj(res))
	case 106: // exp
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == REAL {
			res = math.Exp(val.Value.(float64))
		} else if val.Type == INTGER {
			res = math.Exp(float64(val.Value.(int64)))
		}
		v.OperandStack.Push(makeRealValueObj(res))
	case 107: // abs
		val := v.OperandStack.Pop()
		if val.Type == REAL {
			res := math.Abs(val.Value.(float64))
			v.OperandStack.Push(makeRealValueObj(res))
		} else if val.Type == INTGER {
			res := math.Abs(float64(val.Value.(int64)))
			v.OperandStack.Push(makeIntValueObj(int64(res)))
		}

	case 108: // len of pack
		val := v.OperandStack.Pop()
		if val.Type == PACK {
			length := int64(len(val.Value.(map[PackKey]VMDataObject)))
			v.OperandStack.Push(makeIntValueObj(length))
		} else {
			panic("len() syscall expects a pack type")
		}

	case 109: // asin
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == REAL {
			res = math.Asin(val.Value.(float64))
		} else if val.Type == INTGER {
			res = math.Asin(float64(val.Value.(int64)))
		}
		v.OperandStack.Push(makeRealValueObj(res))

	case 110: // acos
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == REAL {
			res = math.Acos(val.Value.(float64))
		} else if val.Type == INTGER {
			res = math.Acos(float64(val.Value.(int64)))
		}
		v.OperandStack.Push(makeRealValueObj(res))

	case 111: // atan
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == REAL {
			res = math.Atan(val.Value.(float64))
		} else if val.Type == INTGER {
			res = math.Atan(float64(val.Value.(int64)))
		}
		v.OperandStack.Push(makeRealValueObj(res))

	case 112: // atan2
		x_op := v.OperandStack.Pop()
		y_op := v.OperandStack.Pop()
		var x, y float64
		if x_op.Type == REAL {
			x = x_op.Value.(float64)
		} else if x_op.Type == INTGER {
			x = float64(x_op.Value.(int64))
		}
		if y_op.Type == REAL {
			y = y_op.Value.(float64)
		} else if y_op.Type == INTGER {
			y = float64(y_op.Value.(int64))
		}
		res := math.Atan2(y, x)
		v.OperandStack.Push(makeRealValueObj(res))
	case 113: // deg2rad
		val := v.OperandStack.Pop()
		var res float64
		if val.Type == REAL {
			res = val.Value.(float64) * (math.Pi / 180)
		} else if val.Type == INTGER {
			res = float64(val.Value.(int64)) * (math.Pi / 180)
		}
		v.OperandStack.Push(makeRealValueObj(res))
	case 114: // rand_int
		max_op := v.OperandStack.Pop()
		min_op := v.OperandStack.Pop()
		max := max_op.Value.(int64)
		min := min_op.Value.(int64)
		res := min + v.rand.Int63n(max-min+1)
		v.OperandStack.Push(makeIntValueObj(res))
	case 115: // rand_real
		max_op := v.OperandStack.Pop()
		min_op := v.OperandStack.Pop()
		max := max_op.Value.(float64)
		min := min_op.Value.(float64)
		res := min + v.rand.Float64()*(max-min)
		v.OperandStack.Push(makeRealValueObj(res))
	}
}

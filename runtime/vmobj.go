package runtime

import (
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

const (
	INTGER ValueType = iota + 1
	REAL
	STRING
	BOOLEAN
	PACK
	NIL
)

type PackKey struct {
	Type       ValueType
	IntData    int64
	FloatData  float64
	BoolData   bool
	StringData string
}

func (k PackKey) String() string {
	switch k.Type {
	case INTGER:
		return strconv.FormatInt(k.IntData, 10)
	case REAL:
		return strconv.FormatFloat(k.FloatData, 'f', -1, 64)
	case STRING:
		return k.StringData
	default:
		return ""
	}
}

type VMDataObject struct {
	Type ValueType

	IntData    int64
	FloatData  float64
	BoolData   bool
	StringData string
	PackData   map[PackKey]VMDataObject
}

func (d1 VMDataObject) IsEqualTo(d2 VMDataObject) bool {
	if d1.Type != d2.Type {
		return false
	}

	if d1.Type == NIL && d2.Type == NIL {
		return true
	}
	switch d1.Type {
	case INTGER:
		return d1.IntData == d2.IntData
	case REAL:
		return d1.FloatData == d2.FloatData
	case STRING:
		return d1.StringData == d2.StringData
	case BOOLEAN:
		return d1.BoolData == d2.BoolData
	case PACK:
		if d1.PackData == nil || d2.PackData == nil {
			if d1.PackData != nil || d2.PackData == nil {
				return false
			} else if d1.PackData == nil && d2.PackData != nil {
				return false
			}
		}
		if len(d1.PackData) != len(d2.PackData) {
			return false
		}
		for k, v1 := range d1.PackData {
			v2, ok := (d2.PackData)[k]
			if !ok || !v1.IsEqualTo(v2) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func (d1 VMDataObject) IsNotEqualTo(d2 VMDataObject) bool {
	if d1.Type != d2.Type {
		return true
	}

	if d1.Type == NIL && d2.Type == NIL {
		return false
	}
	switch d1.Type {
	case INTGER:
		return d1.IntData != d2.IntData
	case REAL:
		return d1.FloatData != d2.FloatData
	case STRING:
		return d1.StringData != d2.StringData
	case BOOLEAN:
		return d1.BoolData != d2.BoolData
	case PACK:
		if d1.PackData != nil || d2.PackData == nil {
			return false
		} else if d1.PackData == nil && d2.PackData != nil {
			return false
		}
		if len(d1.PackData) != len(d2.PackData) {
			return true
		}
		for k, v1 := range d1.PackData {
			v2, ok := (d2.PackData)[k]
			if !ok || !v1.IsEqualTo(v2) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func (d VMDataObject) String() string {
	switch d.Type {
	case INTGER:
		return strconv.FormatInt(d.IntData, 10)
	case REAL:
		return strconv.FormatFloat(d.FloatData, 'f', -1, 64)
	case STRING:
		return d.StringData
	case BOOLEAN:
		if d.BoolData {
			return "true"
		}
		return "false"
	case PACK:
		if d.PackData == nil {
			return "[]"
		}
		var builder strings.Builder
		builder.WriteString("[")
		i := 0
		for k, v := range d.PackData {
			builder.WriteString(k.String())
			builder.WriteString(": ")
			builder.WriteString(v.String())
			if i < len(d.PackData)-1 {
				builder.WriteString(", ")
			}
			i++
		}
		builder.WriteString("]")
		return builder.String()
	default:
		return "nil"
	}
}

func (r1 VMDataObject) Compare(r2 VMDataObject, floatOp func(float64, float64) bool, intOp func(int64, int64) bool) VMDataObject {
	// If one is REAL, convert both to REAL for comparison
	if r1.Type == REAL || r2.Type == REAL {
		var f1, f2 float64
		if r1.Type == REAL {
			f1 = r1.FloatData
		} else { // r1 is INTGER
			f1 = float64(r1.IntData)
		}
		if r2.Type == REAL {
			f2 = r2.FloatData
		} else { // r2 is INTGER
			f2 = float64(r2.IntData)
		}
		if floatOp != nil {
			return VMDataObject{Type: BOOLEAN, BoolData: floatOp(f1, f2)}
		}
	}

	// Otherwise, both are INTGER
	if r1.Type == INTGER && r2.Type == INTGER {
		if intOp != nil {
			return VMDataObject{Type: BOOLEAN, BoolData: intOp(r1.IntData, r2.IntData)}
		}
	}

	// Default to false for other types or nil ops
	return VMDataObject{Type: BOOLEAN, BoolData: false}
}

func (r1 VMDataObject) Operate(r2 VMDataObject, floatOp func(float64, float64) float64, intOp func(int64, int64) int64, strOp func(string, string) string) VMDataObject {
	switch r1.Type {
	case INTGER:
		switch r2.Type {
		case INTGER:
			if intOp != nil {
				return VMDataObject{Type: INTGER, IntData: intOp(r1.IntData, r2.IntData)}
			}
		case REAL:
			if floatOp != nil {
				return VMDataObject{Type: REAL, FloatData: floatOp(float64(r1.IntData), r2.FloatData)}
			}
		case STRING:
			if strOp != nil {
				return VMDataObject{Type: STRING, StringData: strOp(strconv.FormatInt(r1.IntData, 10), r2.StringData)}
			}
		}
	case REAL:
		switch r2.Type {
		case INTGER:
			if floatOp != nil {
				val1 := r1.FloatData
				val2 := float64(r2.IntData)
				result := floatOp(val1, val2)
				return VMDataObject{Type: REAL, FloatData: result}
			}
		case REAL:
			if floatOp != nil {
				val1 := r1.FloatData
				val2 := r2.FloatData
				result := floatOp(val1, val2)
				return VMDataObject{Type: REAL, FloatData: result}
			}
		case STRING:
			if strOp != nil {
				return VMDataObject{Type: STRING, StringData: strOp(strconv.FormatFloat(r1.FloatData, 'f', -1, 64), r2.StringData)}
			}
		}
	case STRING:
		switch r2.Type {
		case INTGER:
			if strOp != nil {
				return VMDataObject{Type: STRING, StringData: strOp(r1.StringData, strconv.FormatInt(r2.IntData, 10))}
			}
		case REAL:
			if strOp != nil {
				return VMDataObject{Type: STRING, StringData: strOp(r1.StringData, strconv.FormatFloat(r2.FloatData, 'f', -1, 64))}
			}
		case STRING:
			if strOp != nil {
				return VMDataObject{Type: STRING, StringData: strOp(r1.StringData, r2.StringData)}
			}
		}
	}
	spew.Dump(r1, r2)
	panic("Unsupported operation between types" + strconv.FormatInt(int64(r1.Type), 10) + " and " + strconv.FormatInt(int64(r2.Type), 10))
}

func (obj *VMDataObject) CastTo(d_type ValueType) VMDataObject {
	switch d_type {
	case INTGER:
		switch obj.Type {
		case INTGER:
			return *obj
		case REAL:
			val := int64(obj.FloatData)
			return makeIntValueObj(val)
		case STRING:
			val, err := strconv.ParseInt(obj.StringData, 10, 64)
			if err != nil {
				panic("Error Occured in Converting Object - " + err.Error())
			}
			return makeIntValueObj(val)

		default:
			panic("Object cannot be converted to " + string(d_type))

		}

	case REAL:
		switch obj.Type {
		case REAL:
			return *obj
		case INTGER:
			val := float64(obj.IntData)
			return makeRealValueObj(val)
		case STRING:
			val, err := strconv.ParseFloat(obj.StringData, 64)
			if err != nil {
				panic("Error Occured in Converting Object - " + err.Error())
			}
			return makeRealValueObj(val)

		default:
			panic("Object cannot be converted to " + strconv.FormatInt(int64(d_type), 10))

		}

	case STRING:
		switch obj.Type {
		case INTGER:
			return makeStrValueObj(strconv.FormatInt(obj.IntData, 10))
		case REAL:
			return makeStrValueObj(strconv.FormatFloat(obj.FloatData, 'f', -1, 64))

		case BOOLEAN:
			if obj.BoolData {
				return makeStrValueObj("!t")
			} else {
				return makeStrValueObj("!f")
			}

		default:
			panic("Object cannot be converted to " + string(d_type))

		}

	default:
		panic("Object cannot be converted to " + string(d_type))
	}
}

type VMFunctionObject struct {
	JumpPc       int
	IsStandard   bool
	Instructions []VMInstr
}

type VMOp int
type ValueType int

const (
	OpPush VMOp = iota + 1
	OpPop
	OpStoreGlobal
	OpLoadGlobal
	OpStoreLocal
	OpLoadLocal
	OpDefFunc
	OpCall
	OpReturn
	OpSyscall
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpMod
	OpAnd
	OpOr
	OpNot
	OpCmpEq
	OpCmpNeq
	OpCmpGt
	OpCmpLt
	OpCmpGte
	OpCmpLte
	OpJmp
	OpJmpIfFalse
	OpCstInt
	OpCstReal
	OpCstStr
	OpHlt
	OpIndex
	OpMakePack
	OpSetIndex
)

type VMInstr struct {
	Op      VMOp
	Oprand1 VMDataObject
}

func makeIntValueObj(val int64) VMDataObject {
	return VMDataObject{
		Type:    INTGER,
		IntData: val,
	}
}

func makeRealValueObj(val float64) VMDataObject {
	return VMDataObject{
		Type:      REAL,
		FloatData: val,
	}
}

func makeStrValueObj(val string) VMDataObject {
	return VMDataObject{
		Type:       STRING,
		StringData: val,
	}
}

func makeBoolValueObj(val bool) VMDataObject {
	return VMDataObject{
		Type:     BOOLEAN,
		BoolData: val,
	}
}

func makeNilValueObj() VMDataObject {
	return VMDataObject{}
}

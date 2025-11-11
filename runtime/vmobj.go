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
	FUNCTION_ALIAS
)

type PackKey struct {
	Type  ValueType
	Value any
}

func (k PackKey) String() string {
	switch k.Type {
	case INTGER:
		return strconv.FormatInt(k.Value.(int64), 10)
	case REAL:
		return strconv.FormatFloat(k.Value.(float64), 'f', -1, 64)
	case STRING:
		return k.Value.(string)
	default:
		return ""
	}
}

type VMDataObject struct {
	Type  ValueType
	Value any
}

func (d1 VMDataObject) IsEqualTo(d2 VMDataObject) bool {
	if d1.Type != d2.Type {
		return false
	}

	if d1.Type == NIL && d2.Type == NIL {
		return true
	}

	// For PACK type, we need to compare contents
	if d1.Type == PACK {
		p1, ok1 := d1.Value.(map[PackKey]VMDataObject)
		p2, ok2 := d2.Value.(map[PackKey]VMDataObject)
		if !ok1 || !ok2 { // one or both are not maps
			return ok1 == ok2 // true only if both are not maps (e.g. nil)
		}
		if len(p1) != len(p2) {
			return false
		}
		for k, v1 := range p1 {
			v2, ok := p2[k]
			if !ok || !v1.IsEqualTo(v2) {
				return false
			}
		}
		return true
	}

	// For other comparable types
	return d1.Value == d2.Value
}

func (d1 VMDataObject) IsNotEqualTo(d2 VMDataObject) bool {
	return !d1.IsEqualTo(d2) // Simplified logic
}

func (d VMDataObject) String() string {
	if d.Value == nil {
		if d.Type == PACK {
			return "[]"
		}
		return "nil"
	}
	switch d.Type {
	case INTGER:
		return strconv.FormatInt(d.Value.(int64), 10)
	case REAL:
		return strconv.FormatFloat(d.Value.(float64), 'f', -1, 64)
	case STRING:
		return d.Value.(string)
	case BOOLEAN:
		if d.Value.(bool) {
			return "true"
		}
		return "false"
	case PACK:
		packData, ok := d.Value.(map[PackKey]VMDataObject)
		if !ok || packData == nil {
			return "[]"
		}
		var builder strings.Builder
		builder.WriteString("[")
		i := 0
		for k, v := range packData {
			builder.WriteString(k.String())
			builder.WriteString(": ")
			builder.WriteString(v.String())
			if i < len(packData)-1 {
				builder.WriteString(", ")
			}
			i++
		}
		builder.WriteString("]")
		return builder.String()
	case FUNCTION_ALIAS:
		return "func " + d.Value.(string) + ":"
	default:
		return "nil"
	}
}

func (r1 VMDataObject) Compare(r2 VMDataObject, floatOp func(float64, float64) bool, intOp func(int64, int64) bool) VMDataObject {
	var f1, f2 float64
	isFloat := false

	switch r1.Type {
	case REAL:
		f1 = r1.Value.(float64)
		isFloat = true
	case INTGER:
		f1 = float64(r1.Value.(int64))
	default:
		return VMDataObject{Type: BOOLEAN, Value: false}
	}

	switch r2.Type {
	case REAL:
		f2 = r2.Value.(float64)
		isFloat = true
	case INTGER:
		f2 = float64(r2.Value.(int64))
	default:
		return VMDataObject{Type: BOOLEAN, Value: false}
	}

	if isFloat {
		if floatOp != nil {
			return VMDataObject{Type: BOOLEAN, Value: floatOp(f1, f2)}
		}
	} else { // Both are integers
		if intOp != nil {
			return VMDataObject{Type: BOOLEAN, Value: intOp(int64(f1), int64(f2))}
		}
	}

	return VMDataObject{Type: BOOLEAN, Value: false}
}

func (r1 VMDataObject) Operate(r2 VMDataObject, floatOp func(float64, float64) float64, intOp func(int64, int64) int64, strOp func(string, string) string) VMDataObject {
	// String concatenation
	if strOp != nil && (r1.Type == STRING || r2.Type == STRING) {
		return VMDataObject{Type: STRING, Value: strOp(r1.String(), r2.String())}
	}

	var f1, f2 float64
	isFloat := false

	switch r1.Type {
	case REAL:
		f1 = r1.Value.(float64)
		isFloat = true
	case INTGER:
		f1 = float64(r1.Value.(int64))
	default:
		goto unsupported
	}

	switch r2.Type {
	case REAL:
		f2 = r2.Value.(float64)
		isFloat = true
	case INTGER:
		f2 = float64(r2.Value.(int64))
	default:
		goto unsupported
	}

	if isFloat {
		if floatOp != nil {
			return VMDataObject{Type: REAL, Value: floatOp(f1, f2)}
		}
	} else { // Both are integers
		if intOp != nil {
			return VMDataObject{Type: INTGER, Value: intOp(int64(f1), int64(f2))}
		}
	}

unsupported:
	spew.Dump(r1, r2)
	panic("Unsupported operation between types " + r1.String() + " and " + r2.String())
}

func (obj *VMDataObject) CastTo(d_type ValueType) VMDataObject {
	switch d_type {
	case INTGER:
		switch obj.Type {
		case INTGER:
			return *obj
		case REAL:
			return makeIntValueObj(int64(obj.Value.(float64)))
		case STRING:
			val, err := strconv.ParseInt(strings.TrimSpace(obj.Value.(string)), 10, 64)
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
			return makeRealValueObj(float64(obj.Value.(int64)))
		case STRING:
			val, err := strconv.ParseFloat(strings.TrimSpace(obj.Value.(string)), 64)
			if err != nil {
				panic("Error Occured in Converting Object - " + err.Error())
			}
			return makeRealValueObj(val)
		default:
			panic("Object cannot be converted to " + strconv.FormatInt(int64(d_type), 10))
		}
	case STRING:
		return makeStrValueObj(obj.String()) // Simplified using the String() method
	default:
		panic("Object cannot be converted to " + string(d_type))
	}
}

type VMFunctionObject struct {
	JumpPc       int
	IsStandard   bool
	Instructions []VMInstr
}

type VMOp uint8
type ValueType uint8

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
	OpJmpIfEq
	OpJmpIfNeq
	OpJmpIfGt
	OpJmpIfLt
	OpJmpIfGte
	OpJmpIfLte
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
		Type:  INTGER,
		Value: val,
	}
}

func makeRealValueObj(val float64) VMDataObject {
	return VMDataObject{
		Type:  REAL,
		Value: val,
	}
}

func makeStrValueObj(val string) VMDataObject {
	return VMDataObject{
		Type:  STRING,
		Value: val,
	}
}

func makeBoolValueObj(val bool) VMDataObject {
	return VMDataObject{
		Type:  BOOLEAN,
		Value: val,
	}
}

func makeNilValueObj() VMDataObject {
	return VMDataObject{Type: NIL, Value: nil}
}

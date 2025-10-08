package runtime

type CallStack struct {
	stack []int
}

func NewCallStack() *CallStack {
	return &CallStack{stack: make([]int, 0)}
}

func (cs *CallStack) Push(pc int) {
	cs.stack = append(cs.stack, pc)
}

func (cs *CallStack) Pop() int {
	if len(cs.stack) == 0 {
		panic("CallStack underflow")
	}
	val := cs.stack[len(cs.stack)-1]
	cs.stack = cs.stack[:len(cs.stack)-1]
	return val
}

func (cs *CallStack) IsEmpty() bool {
	return len(cs.stack) == 0
}

func (cs *CallStack) GetStack() []int {
	return cs.stack
}

type OperandStack struct {
	stack []VMDataObject
}

func NewOperandStack() *OperandStack {
	return &OperandStack{stack: make([]VMDataObject, 0, 2048)}
}

func (s *OperandStack) Push(obj VMDataObject) {
	s.stack = append(s.stack, obj)
}

func (s *OperandStack) Pop() VMDataObject {
	if len(s.stack) == 0 {
		panic("OperandStack underflow")
	}
	obj := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return obj
}

func (s *OperandStack) Peek() VMDataObject {
	if len(s.stack) == 0 {
		panic("OperandStack is empty")
	}
	return s.stack[len(s.stack)-1]
}

func (s *OperandStack) GetStack() []VMDataObject {
	return s.stack
}

type VMMEMObjectTable struct {
	DataTable      map[string]int
	FunctionTable  map[string]int
	ArrayTable     map[string][]VMDataObject
	DataMemory     []VMDataObject
	FunctionMemory []VMFunctionObject

	currunt_free_dm_pointer int
	currunt_free_fm_pointer int
}

func NewVMMEMObjTable() VMMEMObjectTable {
	return VMMEMObjectTable{
		DataTable:      make(map[string]int),
		FunctionTable:  make(map[string]int),
		ArrayTable:     make(map[string][]VMDataObject),
		DataMemory:     make([]VMDataObject, 0),
		FunctionMemory: make([]VMFunctionObject, 0),

		currunt_free_dm_pointer: 0,
		currunt_free_fm_pointer: 0,
	}
}

func (v *VMMEMObjectTable) MakeObj(name string) {
	v.DataMemory = append(v.DataMemory, VMDataObject{})
	v.DataTable[name] = v.currunt_free_dm_pointer

	v.currunt_free_dm_pointer++
}

func (v *VMMEMObjectTable) GetObj(name string) *VMDataObject {
	idx, ok := v.DataTable[name]
	if !ok {
		panic("VMDataObject not found: " + name)
	}
	return &v.DataMemory[idx]
}

func (v *VMMEMObjectTable) SetObj(name string, data VMDataObject) {
	idx, ok := v.DataTable[name]
	if !ok {
		panic("VMDataObject not found: " + name)
	}
	v.DataMemory[idx] = data
}

func (v *VMMEMObjectTable) HasObj(name string) bool {
	idx, ok := v.DataTable[name]
	if !ok || idx >= len(v.DataMemory) {
		return false
	}
	return true
}

func (v *VMMEMObjectTable) MakeFunc(name string) {
	v.FunctionMemory = append(v.FunctionMemory, VMFunctionObject{})
	v.FunctionTable[name] = v.currunt_free_fm_pointer
	v.currunt_free_fm_pointer++
}

func (v *VMMEMObjectTable) GetFunc(name string) *VMFunctionObject {
	idx, ok := v.FunctionTable[name]
	if !ok || idx >= len(v.FunctionMemory) {
		panic("VMFunctionObject not found: " + name)
	}
	return &v.FunctionMemory[idx]
}

func (v *VMMEMObjectTable) SetFunc(name string, fn VMFunctionObject) {
	idx, ok := v.FunctionTable[name]
	if !ok || idx >= len(v.FunctionMemory) {
		panic("VMFunctionObject not found: " + name)
	}
	v.FunctionMemory[idx] = fn
}

func (v *VMMEMObjectTable) MakeArray(name string) {
	v.ArrayTable[name] = make([]VMDataObject, 0)
}

func (t *VMMEMObjectTable) GetArray(name string) []VMDataObject {
	arr, ok := t.ArrayTable[name]
	if !ok {
		panic("Array not found: " + name)
	}
	return arr
}

func (t *VMMEMObjectTable) PushArrayItem(name string, item VMDataObject) {
	t.ArrayTable[name] = append(t.ArrayTable[name], item)
}

func (t *VMMEMObjectTable) SetArrayItem(name string, idx int, item VMDataObject) {
	if idx >= len(t.ArrayTable[name]) {
		panic("Array index out of bounds: " + name)
	}
	t.ArrayTable[name][idx] = item

}

func (t *VMMEMObjectTable) HasArray(name string) bool {
	_, ok := t.ArrayTable[name]
	return ok
}

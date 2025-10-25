package runtime

type CallStackObject struct {
	PC   int
	Name string
}

type CallStack struct {
	stack []CallStackObject
}

func NewCallStack() *CallStack {
	return &CallStack{stack: make([]CallStackObject, 0)}
}

func (cs *CallStack) Push(pc CallStackObject) {
	cs.stack = append(cs.stack, pc)
}

func (cs *CallStack) Pop() CallStackObject {
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

func (cs *CallStack) GetStack() []CallStackObject {
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

type VMDataObjKey struct {
	Name     string
	ScopeKey string
}

type VMMEMObjectTable struct {
	DataTable      map[VMDataObjKey]int
	FunctionTable  map[string]int
	ArrayTable     map[string][]VMDataObject
	DataMemory     []VMDataObject
	FunctionMemory []VMFunctionObject

	FreeDataMemorySlots     []int
	currunt_free_fm_pointer int
}

func NewVMMEMObjTable() VMMEMObjectTable {
	return VMMEMObjectTable{
		DataTable:      make(map[VMDataObjKey]int),
		FunctionTable:  make(map[string]int),
		ArrayTable:     make(map[string][]VMDataObject),
		DataMemory:     make([]VMDataObject, 0),
		FunctionMemory: make([]VMFunctionObject, 0),

		FreeDataMemorySlots:     make([]int, 0),
		currunt_free_fm_pointer: 0,
	}
}

func (v *VMMEMObjectTable) MakeObj(name string, scopeKey string) {
	key := VMDataObjKey{Name: name, ScopeKey: scopeKey}
	var index int
	if len(v.FreeDataMemorySlots) > 0 {
		// Reuse a free slot
		index = v.FreeDataMemorySlots[len(v.FreeDataMemorySlots)-1]
		v.FreeDataMemorySlots = v.FreeDataMemorySlots[:len(v.FreeDataMemorySlots)-1]
		v.DataMemory[index] = VMDataObject{} // Reset the object
	} else {
		// Allocate a new slot
		index = len(v.DataMemory)
		v.DataMemory = append(v.DataMemory, VMDataObject{})
	}
	v.DataTable[key] = index
}

func (v *VMMEMObjectTable) DeallocateObj(key VMDataObjKey) {
	index, ok := v.DataTable[key]
	if !ok {
		return // Or handle error: trying to deallocate non-existent object
	}

	// Clear the object data
	v.DataMemory[index] = VMDataObject{Type: NIL}

	// Add the index to the free list
	v.FreeDataMemorySlots = append(v.FreeDataMemorySlots, index)

	// Remove from DataTable
	delete(v.DataTable, key)
}

func (v *VMMEMObjectTable) GetObj(name string, scopeKey string) *VMDataObject {
	key := VMDataObjKey{Name: name, ScopeKey: scopeKey}
	idx, ok := v.DataTable[key]
	if !ok {
		panic("VMDataObject not found: " + name)
	}
	return &v.DataMemory[idx]
}

func (v *VMMEMObjectTable) SetObj(name string, data VMDataObject, scopeKey string) {
	key := VMDataObjKey{Name: name, ScopeKey: scopeKey}
	idx, ok := v.DataTable[key]
	if !ok {
		panic("VMDataObject not found: " + name)
	}
	v.DataMemory[idx] = data
}

func (v *VMMEMObjectTable) HasObj(name string, scopeKey string) bool {
	key := VMDataObjKey{Name: name, ScopeKey: scopeKey}
	_, ok := v.DataTable[key]
	return ok
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

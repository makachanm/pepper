package runtime

type CallStackObject struct {
	PC     int
	NameID int
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
	Name     int
	ScopeKey int
}

type VMMEMObjectTable struct {
	DataTable      map[VMDataObjKey]int
	FunctionTable  map[int]int
	ArrayTable     map[int][]VMDataObject
	DataMemory     []VMDataObject
	FunctionMemory []VMFunctionObject

	FreeDataMemorySlots     []int
	currunt_free_fm_pointer int
}

func NewVMMEMObjTable() VMMEMObjectTable {
	return VMMEMObjectTable{
		DataTable:      make(map[VMDataObjKey]int),
		FunctionTable:  make(map[int]int),
		ArrayTable:     make(map[int][]VMDataObject),
		DataMemory:     make([]VMDataObject, 0),
		FunctionMemory: make([]VMFunctionObject, 0),

		FreeDataMemorySlots:     make([]int, 0),
		currunt_free_fm_pointer: 0,
	}
}

func (v *VMMEMObjectTable) MakeObj(nameID int, scopeKeyID int) {
	key := VMDataObjKey{Name: nameID, ScopeKey: scopeKeyID}
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

func (v *VMMEMObjectTable) GetObj(nameID int, scopeKeyID int, vm *VM) *VMDataObject {
	key := VMDataObjKey{Name: nameID, ScopeKey: scopeKeyID}
	idx, ok := v.DataTable[key]
	if !ok {
		panic("VMDataObject not found: " + vm.stringTable[nameID])
	}
	return &v.DataMemory[idx]
}

func (v *VMMEMObjectTable) SetObj(nameID int, data VMDataObject, scopeKeyID int, vm *VM) {
	key := VMDataObjKey{Name: nameID, ScopeKey: scopeKeyID}
	idx, ok := v.DataTable[key]
	if !ok {
		panic("VMDataObject not found: " + vm.stringTable[nameID])
	}
	v.DataMemory[idx] = data
}

func (v *VMMEMObjectTable) HasObj(nameID int, scopeKeyID int, vm *VM) bool {
	key := VMDataObjKey{Name: nameID, ScopeKey: scopeKeyID}
	_, ok := v.DataTable[key]
	return ok
}

func (v *VMMEMObjectTable) MakeFunc(nameID int) {
	v.FunctionMemory = append(v.FunctionMemory, VMFunctionObject{})
	v.FunctionTable[nameID] = v.currunt_free_fm_pointer
	v.currunt_free_fm_pointer++
}

func (v *VMMEMObjectTable) GetFunc(nameID int, vm *VM) *VMFunctionObject {
	idx, ok := v.FunctionTable[nameID]
	if !ok || idx >= len(v.FunctionMemory) {
		panic("VMFunctionObject not found: " + vm.stringTable[nameID])
	}
	return &v.FunctionMemory[idx]
}

func (v *VMMEMObjectTable) SetFunc(nameID int, fn VMFunctionObject, vm *VM) {
	idx, ok := v.FunctionTable[nameID]
	if !ok || idx >= len(v.FunctionMemory) {
		panic("VMFunctionObject not found: " + vm.stringTable[nameID])
	}
	v.FunctionMemory[idx] = fn
}

func (v *VMMEMObjectTable) MakeArray(nameID int) {
	v.ArrayTable[nameID] = make([]VMDataObject, 0)
}

func (t *VMMEMObjectTable) GetArray(nameID int) []VMDataObject {
	arr, ok := t.ArrayTable[nameID]
	if !ok {
		panic("Array not found")
	}
	return arr
}

func (t *VMMEMObjectTable) PushArrayItem(nameID int, item VMDataObject) {
	t.ArrayTable[nameID] = append(t.ArrayTable[nameID], item)
}

func (t *VMMEMObjectTable) SetArrayItem(nameID int, idx int, item VMDataObject) {
	if idx >= len(t.ArrayTable[nameID]) {
		panic("Array index out of bounds")
	}
	t.ArrayTable[nameID][idx] = item

}

func (t *VMMEMObjectTable) HasArray(nameID int) bool {
	_, ok := t.ArrayTable[nameID]
	return ok
}

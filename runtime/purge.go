package runtime

import "sort"

func PurgeVMMEM(mem *VMMEMObjectTable, vm *VM) {
	varsToPurge := vm.functionUsedMemoryVariables[vm.curruntFunctionName]

	indicesToPurge := make([]int, 0, len(varsToPurge))
	for _, name := range varsToPurge {
		if index, exists := mem.DataTable[string(name)]; exists {
			indicesToPurge = append(indicesToPurge, index)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(indicesToPurge)))

	// Remove the elements from DataMemory.
	for _, index := range indicesToPurge {
		// Bounds check for safety
		if index < len(mem.DataMemory) {
			mem.DataMemory = append(mem.DataMemory[:index], mem.DataMemory[index+1:]...)
		}
	}

	// All purged variables must be removed from the lookup table.
	for _, name := range varsToPurge {
		delete(mem.DataTable, string(name))
	}

	mem.currunt_free_dm_pointer -= len(indicesToPurge)

	delete(vm.functionUsedMemoryVariables, vm.curruntFunctionName)

	if len(vm.OperandStack.stack) > 1 {
		newStackSize := len(vm.OperandStack.stack) - vm.callDepth
		if newStackSize < 0 {
			newStackSize = 1
		}
		vm.OperandStack.stack = vm.OperandStack.stack[newStackSize:]
	}

}

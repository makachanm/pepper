package runtime

func PurgeVMMEM(mem *VMMEMObjectTable, vm *VM) {
	returningScope := vm.curruntFunctionName
	if returningScope == "" {
		// Not in a function, nothing to do.
		return
	}

	// Get all objects on the operand stack. These are potential return values.
	operandStack := vm.OperandStack.GetStack()

	// Find all keys in the returning function's scope.
	keysToPurge := []VMDataObjKey{}
	for key := range mem.DataTable {
		if key.ScopeKey == returningScope {
			keysToPurge = append(keysToPurge, key)
		}
	}

	for _, key := range keysToPurge {
		index, ok := mem.DataTable[key]
		if !ok {
			continue
		}
		objToPurge := mem.DataMemory[index]

		isReturnValue := false
		for _, stackObj := range operandStack {
			if objToPurge.IsEqualTo(stackObj) {
				isReturnValue = true
				break
			}
		}

		if !isReturnValue {
			// This object is local to the returned function and is not on the stack.
			// It's safe to deallocate.
			mem.DeallocateObj(key)
		}
	}
}

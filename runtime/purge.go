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

	// Also, perform a safer stack truncation. If the stack is very large,
	// it's good to shrink it to avoid holding onto old return values for too long.
	if len(vm.OperandStack.GetStack()) > 2048 {
		// Keep the top 256 values, which are likely return values from recent functions.
		stack := vm.OperandStack.GetStack()
		vm.OperandStack.stack = stack[len(stack)-256:]
	}
}
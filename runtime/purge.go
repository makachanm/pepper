package runtime

func PurgeVMMEM(mem *VMMEMObjectTable, vm *VM) {
	if len(vm.OperandStack.stack) > 1 {
		newStackSize := len(vm.OperandStack.stack) - vm.callDepth
		if newStackSize < 0 {
			newStackSize = 1
		}
		vm.OperandStack.stack = vm.OperandStack.stack[newStackSize:]
	}

}

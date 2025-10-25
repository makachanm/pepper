package runtime

func PurgeVMMEM(mem *VMMEMObjectTable, vm *VM) {
	if len(vm.OperandStack.GetStack()) > 2048 {
		vm.OperandStack.stack = vm.OperandStack.stack[2048:]

	}
}

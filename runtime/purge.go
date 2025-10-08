package runtime

func PurgeVMMEM(mem *VMMEMObjectTable, vm *VM) {
	if len(vm.OperandStack.GetStack()) > 64 {
		vm.OperandStack.stack = vm.OperandStack.stack[64:]

	}
}

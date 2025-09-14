package compiler

import "pepper/vm"

func (c *Compiler) defineStandardFunctions() {
	c.standardFunctionMaps = map[string][]vm.VMInstr{
		"print": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 1}},
		},
	}
}

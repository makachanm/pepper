package compiler

import "pepper/vm"

func (c *Compiler) defineStandardFunctions() {
	c.standardFunctionMaps = map[string][]vm.VMInstr{
		"print": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 0}},
		},
		"println": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 1}},
		},
		"screen_clear": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 2}},
		},
		"set_source_rgb": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 3}},
		},
		"draw_rect": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 4}},
		},
		"draw_circle": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 5}},
		},
		"draw_line": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 6}},
		},
		"screen_save": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 7}},
		},
		"finish": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 8}},
		},
	}
}

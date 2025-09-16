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
		"set_color": {
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
		"sin": {
			vm.VMInstr{Op: vm.OpSin},
		},
		"cos": {
			vm.VMInstr{Op: vm.OpCos},
		},
		"tan": {
			vm.VMInstr{Op: vm.OpTan},
		},
		"sqrt": {
			vm.VMInstr{Op: vm.OpSqrt},
		},
		"str_len": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 9}},
		},
		"str_sub": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 10}},
		},
		"str_replace": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 11}},
		},
		"http_get": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 12}},
		},
		"http_post": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 13}},
		},
		"http_get_json": {
			vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 14}},
		},
	}
}

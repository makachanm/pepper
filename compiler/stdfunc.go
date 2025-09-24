package compiler

import "pepper/vm"

func (c *Compiler) defineStandardFunctions() {
	c.standardFunctionMaps = map[string][]vm.VMInstr{
		// IO
		"print":      {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 0}}},
		"println":    {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 1}}},
		"readln":     {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 2}}},
		"read_file":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 3}}},
		"write_file": {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 4}}},

		// Math
		"sin":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 100}}},
		"cos":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 101}}},
		"tan":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 102}}},
		"sqrt": {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 103}}},
		"pow":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 104}}},
		"log":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 105}}},
		"exp":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 106}}},
		"abs":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 107}}},
		"len":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 108}}},

		// String
		"str_len":        {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 200}}},
		"str_sub":        {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 201}}},
		"str_replace":    {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 202}}},
		"str_split":      {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 205}}},
		"str_join":       {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 206}}},
		"str_contains":   {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 207}}},
		"str_has_prefix": {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 208}}},
		"str_has_suffix": {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 209}}},
		"str_to_lower":   {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 210}}},
		"str_to_upper":   {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 211}}},
		"str_trim":       {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 212}}},
		"str_index_of":   {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 213}}},
		"json_decode":    {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 501}}},
		"json_encode":    {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 500}}},

		// Graphics
		"gfx_clear":          {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 300}}},
		"gfx_set_source_rgb": {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 301}}},
		"gfx_draw_rect":      {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 302}}},
		"gfx_draw_circle":    {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 303}}},
		"gfx_draw_line":      {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 304}}},
		"gfx_draw_triangle":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 305}}},
		"gfx_draw_bezier":    {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 306}}},
		"gfx_draw_text":      {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 307}}},
		"gfx_save_to_file":   {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 308}}},
		"gfx_finish":         {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 309}}},
		"gfx_set_font_face":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 310}}},
		"gfx_set_font_size":  {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 311}}},

		// HTTP
		"http_get":      {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 400}}},
		"http_post":     {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 401}}},
		"http_get_json": {vm.VMInstr{Op: vm.OpSyscall, Oprand1: vm.VMDataObject{Type: vm.INTGER, IntData: 402}}},
	}
}

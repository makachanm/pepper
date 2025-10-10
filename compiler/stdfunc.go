package compiler

import "pepper/runtime"

func (c *Compiler) defineStandardFunctions() {
	c.standardFunctionMaps = map[string][]runtime.VMInstr{
		"quit": {runtime.VMInstr{Op: runtime.OpHlt, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 0}}},

		"int":  {runtime.VMInstr{Op: runtime.OpCstInt, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 0}}},
		"real": {runtime.VMInstr{Op: runtime.OpCstReal, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 0}}},
		"string":  {runtime.VMInstr{Op: runtime.OpCstStr, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 0}}},

		// IO
		"print":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 0}}},
		"println":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 1}}},
		"readln":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 2}}},
		"read_file":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 3}}},
		"write_file": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 4}}},

		// Math
		"sin":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 100}}},
		"cos":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 101}}},
		"tan":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 102}}},
		"sqrt":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 103}}},
		"pow":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 104}}},
		"log":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 105}}},
		"exp":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 106}}},
		"abs":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 107}}},
		"len":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 108}}},
		"asin":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 109}}},
		"acos":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 110}}},
		"atan":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 111}}},
		"atan2": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 112}}},

		// String
		"strlen":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 200}}},
		"substr":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 201}}},
		"replace":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 202}}},
		"split":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 205}}},
		"join":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 206}}},
		"contains":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 207}}},
		"prefix": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 208}}},
		"suffix": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 209}}},
		"lower":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 210}}},
		"upper":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 211}}},
		"trim":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 212}}},
		"index":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 213}}},
		"json_parse":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 501}}},
		"json_stringify":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 500}}},

		// Graphics
		"clear":          {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 300}}},
		"set_color": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 301}}},
		"draw_rect":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 302}}},
		"draw_circle":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 303}}},
		"draw_line":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 304}}},
		"draw_triangle":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 305}}},
		"draw_bezier":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 306}}},
		"draw_text":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 307}}},
		"save_canvas":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 308}}},
		"render":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 309}}},
		"set_font":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 310}}},
		"set_font_size":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 311}}},
		"wait_event":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 313}}},

		"set_title": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 314}}},
		"resize_window":           {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 315}}},
		"get_width":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 316}}},
		"get_height":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 317}}},

		"set_line_width": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 318}}},
		"stroke":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 319}}},
		"fill":           {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 320}}},
		"path_rect": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 321}}},
		"path_circle":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 322}}},
		"move_to":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 323}}},
		"line_to":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 324}}},
		"close_path":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 325}}},

		"load_sprite":          {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 326}}},
		"create_sprite":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 327}}},
		"destroy_sprite":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 328}}},
		"draw_sprite":          {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 329}}},
		"set_sprite_rotation":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 330}}},
		"set_sprite_scale":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 331}}},

		// HTTP
		"http_get":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 400}}},
		"http_post":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 401}}},
		"http_get_json": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 402}}},

		// Time
		"sleep": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 600}}},
	}
}

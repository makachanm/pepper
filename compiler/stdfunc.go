package compiler

import "pepper/runtime"

func (c *Compiler) defineStandardFunctions() {
	c.standardFunctionMaps = map[string][]runtime.VMInstr{
		"quit": {runtime.VMInstr{Op: runtime.OpHlt, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 0}}},

		"to_int":  {runtime.VMInstr{Op: runtime.OpCstInt, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 0}}},
		"to_real": {runtime.VMInstr{Op: runtime.OpCstReal, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 0}}},
		"to_str":  {runtime.VMInstr{Op: runtime.OpCstStr, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 0}}},

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
		"2atan": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 112}}},

		// String
		"str_len":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 200}}},
		"str_sub":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 201}}},
		"str_replace":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 202}}},
		"str_split":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 205}}},
		"str_join":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 206}}},
		"str_contains":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 207}}},
		"str_has_prefix": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 208}}},
		"str_has_suffix": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 209}}},
		"str_to_lower":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 210}}},
		"str_to_upper":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 211}}},
		"str_trim":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 212}}},
		"str_index_of":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 213}}},
		"json_decode":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 501}}},
		"json_encode":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 500}}},

		// Graphics
		"gfx_clear":          {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 300}}},
		"gfx_set_source_rgb": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 301}}},
		"gfx_draw_rect":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 302}}},
		"gfx_draw_circle":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 303}}},
		"gfx_draw_line":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 304}}},
		"gfx_draw_triangle":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 305}}},
		"gfx_draw_bezier":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 306}}},
		"gfx_draw_text":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 307}}},
		"gfx_save_to_file":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 308}}},
		"gfx_finish":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 309}}},
		"gfx_set_font_face":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 310}}},
		"gfx_set_font_size":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 311}}},
		"gfx_wait_event":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 313}}},

		"gfx_set_window_title": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 314}}},
		"gfx_resize":           {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 315}}},
		"gfx_get_width":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 316}}},
		"gfx_get_height":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 317}}},

		"gfx_set_line_width": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 318}}},
		"gfx_stroke":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 319}}},
		"gfx_fill":           {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 320}}},
		"gfx_path_rectangle": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 321}}},
		"gfx_path_circle":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 322}}},
		"gfx_path_move_to":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 323}}},
		"gfx_path_line_to":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 324}}},
		"gfx_path_close":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 325}}},

		// HTTP
		"http_get":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 400}}},
		"http_post":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 401}}},
		"http_get_json": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 402}}},

		// Time
		"sleep": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, IntData: 600}}},
	}
}

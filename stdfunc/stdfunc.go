package stdfunc

import "pepper/runtime"

func DefineStandardFunctions() map[string][]runtime.VMInstr {
	return map[string][]runtime.VMInstr{
		"quit": {runtime.VMInstr{Op: runtime.OpHlt, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(0)}}},

		"int":    {runtime.VMInstr{Op: runtime.OpCstInt, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(0)}}},
		"real":   {runtime.VMInstr{Op: runtime.OpCstReal, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(0)}}},
		"string": {runtime.VMInstr{Op: runtime.OpCstStr, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(0)}}},

		// IO
		"print":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(0)}}},
		"println":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(1)}}},
		"readln":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(2)}}},
		"read_file":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(3)}}},
		"write_file": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(4)}}},

		// Math
		"sin":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(100)}}},
		"cos":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(101)}}},
		"tan":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(102)}}},
		"sqrt":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(103)}}},
		"pow":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(104)}}},
		"log":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(105)}}},
		"exp":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(106)}}},
		"abs":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(107)}}},
		"len":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(108)}}},
		"asin":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(109)}}},
		"acos":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(110)}}},
		"atan":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(111)}}},
		"atan2":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(112)}}},
		"deg2rad":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(113)}}},
		"rand_int":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(114)}}},
		"rand_real": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(115)}}},

		// String
		"strlen":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(200)}}},
		"substr":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(201)}}},
		"replace":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(202)}}},
		"split":          {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(205)}}},
		"join":           {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(206)}}},
		"contains":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(207)}}},
		"prefix":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(208)}}},
		"suffix":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(209)}}},
		"lower":          {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(210)}}},
		"upper":          {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(211)}}},
		"trim":           {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(212)}}},
		"index":          {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(213)}}},
		"json_parse":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(501)}}},
		"json_stringify": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(500)}}},

		// Graphics
		"clear":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(300)}}},
		"set_color":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(301)}}},
		"draw_rect":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(302)}}},
		"draw_dot":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(332)}}},
		"draw_circle":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(303)}}},
		"draw_line":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(304)}}},
		"draw_triangle": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(305)}}},
		"draw_bezier":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(306)}}},
		"draw_text":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(307)}}},
		"save_canvas":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(308)}}},
		"render":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(309)}}},
		"set_font":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(310)}}},
		"set_font_size": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(311)}}},
		"wait_event":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(313)}}},

		"set_title":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(314)}}},
		"resize_window": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(315)}}},
		"get_width":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(316)}}},
		"get_height":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(317)}}},

		"set_line_width": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(318)}}},
		"stroke":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(319)}}},
		"fill":           {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(320)}}},
		"path_rect":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(321)}}},
		"path_circle":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(322)}}},
		"move_to":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(323)}}},
		"line_to":        {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(324)}}},
		"close_path":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(325)}}},

		"load_sprite":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(326)}}},
		"create_sprite":       {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(327)}}},
		"destroy_sprite":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(328)}}},
		"draw_sprite":         {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(329)}}},
		"set_sprite_rotation": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(330)}}},
		"set_sprite_scale":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(331)}}},
		"set_rgba":            {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(312)}}},
		"set_mask":            {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(333)}}},
		"reset_mask":          {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(334)}}},

		// HTTP
		"http_get":      {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(400)}}},
		"http_post":     {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(401)}}},
		"http_get_json": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(402)}}},

		// Time
		"sleep":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(600)}}},
		"now":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(601)}}},
		"year":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(602)}}},
		"month":  {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(603)}}},
		"day":    {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(604)}}},
		"hour":   {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(605)}}},
		"minute": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(606)}}},
		"second": {runtime.VMInstr{Op: runtime.OpSyscall, Oprand1: runtime.VMDataObject{Type: runtime.INTGER, Value: int64(607)}}},
	}
}

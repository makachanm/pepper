package runtime

func doSyscallGfx(v VM, code int64) {
	switch code {
	case 2: // gfx_clear
		GfxClear(v.OperandStack)
	case 3: // gfx_set_source_rgb
		GfxSetSourceRGB(v.OperandStack)
	case 4: // gfx_draw_rect
		GfxDrawRect(v.OperandStack)
	case 5: // gfx_draw_circle
		GfxDrawCircle(v.OperandStack)
	case 6: // gfx_draw_line
		GfxDrawLine(v.OperandStack)
	case 7: // gfx_save_to_file
		GfxSaveToFile(v.OperandStack)
	case 8: // gfx_finish
		GfxFinish()
	}
}

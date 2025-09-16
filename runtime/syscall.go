package runtime

import (
	"fmt"
)

func doSyscall(vm VM, code int64) {
	switch code {
	case 0: // Print top of stack
		val := vm.OperandStack.Pop()
		fmt.Print(val.String())
	case 1: // Print and pop top of stack
		val := vm.OperandStack.Pop()
		fmt.Printf("%s\n", val.String())
	case 2: // gfx_clear
		GfxClear(vm.OperandStack)
	case 3: // gfx_set_source_rgb
		GfxSetSourceRGB(vm.OperandStack)
	case 4: // gfx_draw_rect
		GfxDrawRect(vm.OperandStack)
	case 5: // gfx_draw_circle
		GfxDrawCircle(vm.OperandStack)
	case 6: // gfx_draw_line
		GfxDrawLine(vm.OperandStack)
	case 7: // gfx_save_to_file
		GfxSaveToFile(vm.OperandStack)
	case 8: // gfx_finish
		GfxFinish()
	default:
		panic("Unknown syscall code")
	}
}

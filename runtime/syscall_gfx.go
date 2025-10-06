package runtime

import (
	"pepper/vm"
)

func doSyscallGfx(vmInstance VM, code int64) {
	switch code {
	case 300: // gfx_clear
		GfxClear(vmInstance.OperandStack)
	case 301: // gfx_set_source_rgb
		GfxSetSourceRGB(vmInstance.OperandStack)
	case 302: // gfx_draw_rect
		GfxDrawRect(vmInstance.OperandStack)
	case 303: // gfx_draw_circle
		GfxDrawCircle(vmInstance.OperandStack)
	case 304: // gfx_draw_line
		GfxDrawLine(vmInstance.OperandStack)
	case 305: // gfx_draw_triangle
		GfxDrawTriangle(vmInstance.OperandStack)
	case 306: // gfx_draw_bezier
		GfxDrawBezier(vmInstance.OperandStack)
	case 307: // gfx_draw_text
		GfxDrawText(vmInstance.OperandStack)
	case 308: // gfx_save_to_file
		GfxSaveToFile(vmInstance.OperandStack)
	case 309: // gfx_finish
		GfxFinish()
	case 310: // gfx_set_font_face
		GfxSetFontFace(vmInstance.OperandStack)
	case 311: // gfx_set_font_size
		GfxSetFontSize(vmInstance.OperandStack)

	case 313: // gfx_wait_event
		GfxWaitEvent(vmInstance.OperandStack)
	}
}

func eventToPack(event Event) vm.VMDataObject {
	pack := make(map[vm.PackKey]vm.VMDataObject)
	pack[vm.PackKey{Type: vm.STRING, StringData: "type"}] = vm.VMDataObject{Type: vm.STRING, StringData: string(event.Type)}

	switch event.Type {
	case EventTypeMouseMotion, EventTypeMouseButtonDown, EventTypeMouseButtonUp:
		pack[vm.PackKey{Type: vm.STRING, StringData: "x"}] = vm.VMDataObject{Type: vm.INTGER, IntData: int64(event.X)}
		pack[vm.PackKey{Type: vm.STRING, StringData: "y"}] = vm.VMDataObject{Type: vm.INTGER, IntData: int64(event.Y)}
	}

	if event.Type == EventTypeMouseButtonDown || event.Type == EventTypeMouseButtonUp {
		pack[vm.PackKey{Type: vm.STRING, StringData: "button"}] = vm.VMDataObject{Type: vm.INTGER, IntData: int64(event.Button)}
	}

	if event.Type == EventTypeKeyDown || event.Type == EventTypeKeyUp {
		pack[vm.PackKey{Type: vm.STRING, StringData: "key_name"}] = vm.VMDataObject{Type: vm.STRING, StringData: event.KeyName}
	}

	return vm.VMDataObject{Type: vm.PACK, PackData: pack}
}

func GfxWaitEvent(stack *vm.OperandStack) {
	event := EventQueue.Dequeue()
	stack.Push(eventToPack(event))
}
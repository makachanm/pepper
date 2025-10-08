package runtime

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

	case 314:
		GfxSetWindowTitle(vmInstance.OperandStack)
	case 315:
		GfxResize(vmInstance.OperandStack)
	case 316:
		GfxGetWidth(vmInstance.OperandStack)
	case 317:
		GfxGetHeight(vmInstance.OperandStack)

	case 318:
		GfxSetLineWidth(vmInstance.OperandStack)
	case 319:
		GfxStroke(vmInstance.OperandStack)
	case 320:
		GfxFill(vmInstance.OperandStack)
	case 321:
		GfxPathRectangle(vmInstance.OperandStack)
	case 322:
		GfxPathCircle(vmInstance.OperandStack)
	case 323:
		GfxPathMoveTo(vmInstance.OperandStack)
	case 324:
		GfxPathLineTo(vmInstance.OperandStack)
	case 325:
		GfxPathClose(vmInstance.OperandStack)
	case 326:
		GfxLoadSprite(vmInstance.OperandStack)
	case 327:
		GfxCreateSprite(vmInstance.OperandStack)
	case 328:
		GfxDestroySprite(vmInstance.OperandStack)
	case 329:
		GfxDrawSprite(vmInstance.OperandStack)
	case 330:
		GfxSetSpriteRotation(vmInstance.OperandStack)
	case 331:
		GfxSetSpriteScale(vmInstance.OperandStack)
	}
}

func eventToPack(event Event) VMDataObject {
	pack := make(map[PackKey]VMDataObject)
	pack[PackKey{Type: STRING, StringData: "type"}] = VMDataObject{Type: STRING, StringData: string(event.Type)}

	switch event.Type {
	case EventTypeMouseMotion, EventTypeMouseButtonDown, EventTypeMouseButtonUp:
		pack[PackKey{Type: STRING, StringData: "x"}] = VMDataObject{Type: INTGER, IntData: int64(event.X)}
		pack[PackKey{Type: STRING, StringData: "y"}] = VMDataObject{Type: INTGER, IntData: int64(event.Y)}
	}

	if event.Type == EventTypeMouseButtonDown || event.Type == EventTypeMouseButtonUp {
		pack[PackKey{Type: STRING, StringData: "button"}] = VMDataObject{Type: INTGER, IntData: int64(event.Button)}
	}

	if event.Type == EventTypeKeyDown || event.Type == EventTypeKeyUp {
		pack[PackKey{Type: STRING, StringData: "key_name"}] = VMDataObject{Type: STRING, StringData: event.KeyName}
	}

	return VMDataObject{Type: PACK, PackData: pack}
}

func GfxWaitEvent(stack *OperandStack) {
	if event, ok := EventQueue.DequeueNonBlocking(); ok {
		stack.Push(eventToPack(event))
	} else {
		stack.Push(VMDataObject{Type: NIL})
	}
}

package runtime

import (
	"pepper/vm"
)

var Gfx *CairoGraphics

func GfxNew(width, height int) {
	Gfx = NewCairoGraphics(width, height)
}

func GfxClear(stack *vm.OperandStack) {
	Gfx.Clear()
}

func GfxSetSourceRGB(stack *vm.OperandStack) {
	b := stack.Pop().FloatData
	g := stack.Pop().FloatData
	r := stack.Pop().FloatData
	Gfx.SetSourceRGB(r, g, b)
}

func GfxDrawRect(stack *vm.OperandStack) {
	height := stack.Pop().IntData
	width := stack.Pop().IntData
	y := stack.Pop().IntData
	x := stack.Pop().IntData
	Gfx.DrawRect(int(x), int(y), int(width), int(height))
}

func GfxDrawCircle(stack *vm.OperandStack) {
	radius := stack.Pop().IntData
	y := stack.Pop().IntData
	x := stack.Pop().IntData
	Gfx.DrawCircle(int(x), int(y), int(radius))
}

func GfxDrawLine(stack *vm.OperandStack) {
	y2 := stack.Pop().IntData
	x2 := stack.Pop().IntData
	y1 := stack.Pop().IntData
	x1 := stack.Pop().IntData
	Gfx.DrawLine(int(x1), int(y1), int(x2), int(y2))
}

func GfxSaveToFile(stack *vm.OperandStack) {
	filename := stack.Pop().StringData
	Gfx.SaveToFile(filename)
}

func GfxFinish() {
	Gfx.Finish()
}

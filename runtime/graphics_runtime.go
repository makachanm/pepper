package runtime

import (
	"pepper/vm"
	"sync"
)

var Gfx Graphics

func GfxNew(width, height int, wg *sync.WaitGroup) {
	Gfx = NewGraphics(width, height, wg)
}

func GfxClear(stack *vm.OperandStack) {
	Gfx.Clear()
}

func GfxSetSourceRGB(stack *vm.OperandStack) {
	b := stack.Pop().IntData
	g := stack.Pop().IntData
	r := stack.Pop().IntData
	Gfx.SetSourceRGB(float64(r), float64(g), float64(b))
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

func GfxDrawTriangle(stack *vm.OperandStack) {
	y3 := stack.Pop().IntData
	x3 := stack.Pop().IntData
	y2 := stack.Pop().IntData
	x2 := stack.Pop().IntData
	y1 := stack.Pop().IntData
	x1 := stack.Pop().IntData
	Gfx.DrawTriangle(int(x1), int(y1), int(x2), int(y2), int(x3), int(y3))
}

func GfxDrawBezier(stack *vm.OperandStack) {
	y4 := stack.Pop().IntData
	x4 := stack.Pop().IntData
	y3 := stack.Pop().IntData
	x3 := stack.Pop().IntData
	y2 := stack.Pop().IntData
	x2 := stack.Pop().IntData
	y1 := stack.Pop().IntData
	x1 := stack.Pop().IntData
	Gfx.DrawBezier(int(x1), int(y1), int(x2), int(y2), int(x3), int(y3), int(x4), int(y4))
}

func GfxDrawText(stack *vm.OperandStack) {
	text := stack.Pop().StringData
	y := stack.Pop().IntData
	x := stack.Pop().IntData
	Gfx.DrawText(int(x), int(y), text)
}

func GfxSetFontFace(stack *vm.OperandStack) {
	fontFace := stack.Pop().StringData
	Gfx.SetFontFace(fontFace)
}

func GfxSetFontSize(stack *vm.OperandStack) {
	size := stack.Pop().IntData
	Gfx.SetFontSize(float64(size))
}

func GfxSaveToFile(stack *vm.OperandStack) {
	filename := stack.Pop().StringData
	Gfx.SaveToFile(filename)
}

func GfxFinish() {
	Gfx.Finish()
}

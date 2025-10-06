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
	bObj := stack.Pop()
	gObj := stack.Pop()
	rObj := stack.Pop()

	var r, g, b float64

	if rObj.Type == vm.REAL {
		r = rObj.FloatData
	} else {
		r = float64(rObj.IntData) / 255.0
	}

	if gObj.Type == vm.REAL {
		g = gObj.FloatData
	} else {
		g = float64(gObj.IntData) / 255.0
	}

	if bObj.Type == vm.REAL {
		b = bObj.FloatData
	} else {
		b = float64(bObj.IntData) / 255.0
	}

	Gfx.SetSourceRGB(r, g, b)
}

func GfxDrawRect(stack *vm.OperandStack) {
	heightObj := stack.Pop()
	widthObj := stack.Pop()
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y, width, height int

	if xObj.Type == vm.REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == vm.REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	if widthObj.Type == vm.REAL {
		width = int(widthObj.FloatData)
	} else {
		width = int(widthObj.IntData)
	}

	if heightObj.Type == vm.REAL {
		height = int(heightObj.FloatData)
	} else {
		height = int(heightObj.IntData)
	}

	Gfx.DrawRect(x, y, width, height)
}

func GfxDrawCircle(stack *vm.OperandStack) {
	radiusObj := stack.Pop()
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y, radius int

	if xObj.Type == vm.REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == vm.REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	if radiusObj.Type == vm.REAL {
		radius = int(radiusObj.FloatData)
	} else {
		radius = int(radiusObj.IntData)
	}

	Gfx.DrawCircle(x, y, radius)
}

func GfxDrawLine(stack *vm.OperandStack) {
	y2Obj := stack.Pop()
	x2Obj := stack.Pop()
	y1Obj := stack.Pop()
	x1Obj := stack.Pop()

	var x1, y1, x2, y2 int

	if x1Obj.Type == vm.REAL {
		x1 = int(x1Obj.FloatData)
	} else {
		x1 = int(x1Obj.IntData)
	}

	if y1Obj.Type == vm.REAL {
		y1 = int(y1Obj.FloatData)
	} else {
		y1 = int(y1Obj.IntData)
	}

	if x2Obj.Type == vm.REAL {
		x2 = int(x2Obj.FloatData)
	} else {
		x2 = int(x2Obj.IntData)
	}

	if y2Obj.Type == vm.REAL {
		y2 = int(y2Obj.FloatData)
	} else {
		y2 = int(y2Obj.IntData)
	}

	Gfx.DrawLine(x1, y1, x2, y2)
}

func GfxDrawTriangle(stack *vm.OperandStack) {
	y3Obj := stack.Pop()
	x3Obj := stack.Pop()
	y2Obj := stack.Pop()
	x2Obj := stack.Pop()
	y1Obj := stack.Pop()
	x1Obj := stack.Pop()

	var x1, y1, x2, y2, x3, y3 int

	if x1Obj.Type == vm.REAL {
		x1 = int(x1Obj.FloatData)
	} else {
		x1 = int(x1Obj.IntData)
	}

	if y1Obj.Type == vm.REAL {
		y1 = int(y1Obj.FloatData)
	} else {
		y1 = int(y1Obj.IntData)
	}

	if x2Obj.Type == vm.REAL {
		x2 = int(x2Obj.FloatData)
	} else {
		x2 = int(x2Obj.IntData)
	}

	if y2Obj.Type == vm.REAL {
		y2 = int(y2Obj.FloatData)
	} else {
		y2 = int(y2Obj.IntData)
	}

	if x3Obj.Type == vm.REAL {
		x3 = int(x3Obj.FloatData)
	} else {
		x3 = int(x3Obj.IntData)
	}

	if y3Obj.Type == vm.REAL {
		y3 = int(y3Obj.FloatData)
	} else {
		y3 = int(y3Obj.IntData)
	}

	Gfx.DrawTriangle(x1, y1, x2, y2, x3, y3)
}

func GfxDrawBezier(stack *vm.OperandStack) {
	y4Obj := stack.Pop()
	x4Obj := stack.Pop()
	y3Obj := stack.Pop()
	x3Obj := stack.Pop()
	y2Obj := stack.Pop()
	x2Obj := stack.Pop()
	y1Obj := stack.Pop()
	x1Obj := stack.Pop()

	var x1, y1, x2, y2, x3, y3, x4, y4 int

	if x1Obj.Type == vm.REAL {
		x1 = int(x1Obj.FloatData)
	} else {
		x1 = int(x1Obj.IntData)
	}

	if y1Obj.Type == vm.REAL {
		y1 = int(y1Obj.FloatData)
	} else {
		y1 = int(y1Obj.IntData)
	}

	if x2Obj.Type == vm.REAL {
		x2 = int(x2Obj.FloatData)
	} else {
		x2 = int(x2Obj.IntData)
	}

	if y2Obj.Type == vm.REAL {
		y2 = int(y2Obj.FloatData)
	} else {
		y2 = int(y2Obj.IntData)
	}

	if x3Obj.Type == vm.REAL {
		x3 = int(x3Obj.FloatData)
	} else {
		x3 = int(x3Obj.IntData)
	}

	if y3Obj.Type == vm.REAL {
		y3 = int(y3Obj.FloatData)
	} else {
		y3 = int(y3Obj.IntData)
	}

	if x4Obj.Type == vm.REAL {
		x4 = int(x4Obj.FloatData)
	} else {
		x4 = int(x4Obj.IntData)
	}

	if y4Obj.Type == vm.REAL {
		y4 = int(y4Obj.FloatData)
	} else {
		y4 = int(y4Obj.IntData)
	}

	Gfx.DrawBezier(x1, y1, x2, y2, x3, y3, x4, y4)
}

func GfxDrawText(stack *vm.OperandStack) {
	text := stack.Pop().StringData
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y int

	if xObj.Type == vm.REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == vm.REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	Gfx.DrawText(x, y, text)
}

func GfxSetFontFace(stack *vm.OperandStack) {
	fontFace := stack.Pop().StringData
	Gfx.SetFontFace(fontFace)
}

func GfxSetFontSize(stack *vm.OperandStack) {
	sizeObj := stack.Pop()
	var size float64
	if sizeObj.Type == vm.REAL {
		size = sizeObj.FloatData
	} else {
		size = float64(sizeObj.IntData)
	}
	Gfx.SetFontSize(size)
}

func GfxSaveToFile(stack *vm.OperandStack) {
	filename := stack.Pop().StringData
	Gfx.SaveToFile(filename)
}

func GfxFinish() {
	Gfx.Finish()
}

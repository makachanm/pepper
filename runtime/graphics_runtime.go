package runtime

import (
	"sync"
)

var Gfx Graphics

func GfxNew(width, height int, wg *sync.WaitGroup) {
	Gfx = NewGraphics(width, height, wg)
}

func GfxResize(stack *OperandStack) {
	heightObj := stack.Pop()
	widthObj := stack.Pop()

	var width, height int

	if widthObj.Type == REAL {
		width = int(widthObj.FloatData)
	} else {
		width = int(widthObj.IntData)
	}

	if heightObj.Type == REAL {
		height = int(heightObj.FloatData)
	} else {
		height = int(heightObj.IntData)
	}

	Gfx.Resize(width, height)
}

func GfxGetWidth(stack *OperandStack) {
	width, _ := Gfx.GetDimensions()
	stack.Push(VMDataObject{Type: INTGER, IntData: int64(width)})
}

func GfxGetHeight(stack *OperandStack) {
	_, height := Gfx.GetDimensions()
	stack.Push(VMDataObject{Type: INTGER, IntData: int64(height)})
}

func GfxSetWindowTitle(stack *OperandStack) {
	title := stack.Pop().StringData
	Gfx.SetWindowTitle(title)
}

func GfxClear(stack *OperandStack) {
	Gfx.Clear()
}

func GfxSetSourceRGB(stack *OperandStack) {
	bObj := stack.Pop()
	gObj := stack.Pop()
	rObj := stack.Pop()

	var r, g, b float64

	if rObj.Type == REAL {
		r = rObj.FloatData
	} else {
		r = float64(rObj.IntData) / 255.0
	}

	if gObj.Type == REAL {
		g = gObj.FloatData
	} else {
		g = float64(gObj.IntData) / 255.0
	}

	if bObj.Type == REAL {
		b = bObj.FloatData
	} else {
		b = float64(bObj.IntData) / 255.0
	}

	Gfx.SetSourceRGB(r, g, b)
}

func GfxDrawRect(stack *OperandStack) {
	heightObj := stack.Pop()
	widthObj := stack.Pop()
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y, width, height int

	if xObj.Type == REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	if widthObj.Type == REAL {
		width = int(widthObj.FloatData)
	} else {
		width = int(widthObj.IntData)
	}

	if heightObj.Type == REAL {
		height = int(heightObj.FloatData)
	} else {
		height = int(heightObj.IntData)
	}

	Gfx.DrawRect(x, y, width, height)
}

func GfxDrawCircle(stack *OperandStack) {
	radiusObj := stack.Pop()
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y, radius int

	if xObj.Type == REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	if radiusObj.Type == REAL {
		radius = int(radiusObj.FloatData)
	} else {
		radius = int(radiusObj.IntData)
	}

	Gfx.DrawCircle(x, y, radius)
}

func GfxDrawLine(stack *OperandStack) {
	y2Obj := stack.Pop()
	x2Obj := stack.Pop()
	y1Obj := stack.Pop()
	x1Obj := stack.Pop()

	var x1, y1, x2, y2 int

	if x1Obj.Type == REAL {
		x1 = int(x1Obj.FloatData)
	} else {
		x1 = int(x1Obj.IntData)
	}

	if y1Obj.Type == REAL {
		y1 = int(y1Obj.FloatData)
	} else {
		y1 = int(y1Obj.IntData)
	}

	if x2Obj.Type == REAL {
		x2 = int(x2Obj.FloatData)
	} else {
		x2 = int(x2Obj.IntData)
	}

	if y2Obj.Type == REAL {
		y2 = int(y2Obj.FloatData)
	} else {
		y2 = int(y2Obj.IntData)
	}

	Gfx.DrawLine(x1, y1, x2, y2)
}

func GfxDrawTriangle(stack *OperandStack) {
	y3Obj := stack.Pop()
	x3Obj := stack.Pop()
	y2Obj := stack.Pop()
	x2Obj := stack.Pop()
	y1Obj := stack.Pop()
	x1Obj := stack.Pop()

	var x1, y1, x2, y2, x3, y3 int

	if x1Obj.Type == REAL {
		x1 = int(x1Obj.FloatData)
	} else {
		x1 = int(x1Obj.IntData)
	}

	if y1Obj.Type == REAL {
		y1 = int(y1Obj.FloatData)
	} else {
		y1 = int(y1Obj.IntData)
	}

	if x2Obj.Type == REAL {
		x2 = int(x2Obj.FloatData)
	} else {
		x2 = int(x2Obj.IntData)
	}

	if y2Obj.Type == REAL {
		y2 = int(y2Obj.FloatData)
	} else {
		y2 = int(y2Obj.IntData)
	}

	if x3Obj.Type == REAL {
		x3 = int(x3Obj.FloatData)
	} else {
		x3 = int(x3Obj.IntData)
	}

	if y3Obj.Type == REAL {
		y3 = int(y3Obj.FloatData)
	} else {
		y3 = int(y3Obj.IntData)
	}

	Gfx.DrawTriangle(x1, y1, x2, y2, x3, y3)
}

func GfxDrawBezier(stack *OperandStack) {
	y4Obj := stack.Pop()
	x4Obj := stack.Pop()
	y3Obj := stack.Pop()
	x3Obj := stack.Pop()
	y2Obj := stack.Pop()
	x2Obj := stack.Pop()
	y1Obj := stack.Pop()
	x1Obj := stack.Pop()

	var x1, y1, x2, y2, x3, y3, x4, y4 int

	if x1Obj.Type == REAL {
		x1 = int(x1Obj.FloatData)
	} else {
		x1 = int(x1Obj.IntData)
	}

	if y1Obj.Type == REAL {
		y1 = int(y1Obj.FloatData)
	} else {
		y1 = int(y1Obj.IntData)
	}

	if x2Obj.Type == REAL {
		x2 = int(x2Obj.FloatData)
	} else {
		x2 = int(x2Obj.IntData)
	}

	if y2Obj.Type == REAL {
		y2 = int(y2Obj.FloatData)
	} else {
		y2 = int(y2Obj.IntData)
	}

	if x3Obj.Type == REAL {
		x3 = int(x3Obj.FloatData)
	} else {
		x3 = int(x3Obj.IntData)
	}

	if y3Obj.Type == REAL {
		y3 = int(y3Obj.FloatData)
	} else {
		y3 = int(y3Obj.IntData)
	}

	if x4Obj.Type == REAL {
		x4 = int(x4Obj.FloatData)
	} else {
		x4 = int(x4Obj.IntData)
	}

	if y4Obj.Type == REAL {
		y4 = int(y4Obj.FloatData)
	} else {
		y4 = int(y4Obj.IntData)
	}

	Gfx.DrawBezier(x1, y1, x2, y2, x3, y3, x4, y4)
}

func GfxDrawText(stack *OperandStack) {
	text := stack.Pop().StringData
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y int

	if xObj.Type == REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	Gfx.DrawText(x, y, text)
}

func GfxSetFontFace(stack *OperandStack) {
	fontFace := stack.Pop().StringData
	Gfx.SetFontFace(fontFace)
}

func GfxSetFontSize(stack *OperandStack) {
	sizeObj := stack.Pop()
	var size float64
	if sizeObj.Type == REAL {
		size = sizeObj.FloatData
	} else {
		size = float64(sizeObj.IntData)
	}
	Gfx.SetFontSize(size)
}

func GfxSaveToFile(stack *OperandStack) {
	filename := stack.Pop().StringData
	Gfx.SaveToFile(filename)
}

func GfxFinish() {
	Gfx.Finish()
}

func GfxSetLineWidth(stack *OperandStack) {
	widthObj := stack.Pop()
	var width float64
	if widthObj.Type == REAL {
		width = widthObj.FloatData
	} else {
		width = float64(widthObj.IntData)
	}
	Gfx.SetLineWidth(width)
}

func GfxStroke(stack *OperandStack) {
	Gfx.Stroke()
}

func GfxFill(stack *OperandStack) {
	Gfx.Fill()
}

func GfxPathRectangle(stack *OperandStack) {
	heightObj := stack.Pop()
	widthObj := stack.Pop()
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y, width, height int

	if xObj.Type == REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	if widthObj.Type == REAL {
		width = int(widthObj.FloatData)
	} else {
		width = int(widthObj.IntData)
	}

	if heightObj.Type == REAL {
		height = int(heightObj.FloatData)
	} else {
		height = int(heightObj.IntData)
	}

	Gfx.PathRectangle(x, y, width, height)
}

func GfxPathCircle(stack *OperandStack) {
	radiusObj := stack.Pop()
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y, radius int

	if xObj.Type == REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	if radiusObj.Type == REAL {
		radius = int(radiusObj.FloatData)
	} else {
		radius = int(radiusObj.IntData)
	}

	Gfx.PathCircle(x, y, radius)
}

func GfxPathMoveTo(stack *OperandStack) {
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y int

	if xObj.Type == REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	Gfx.PathMoveTo(x, y)
}

func GfxPathLineTo(stack *OperandStack) {
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y int

	if xObj.Type == REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	Gfx.PathLineTo(x, y)
}

func GfxPathClose(stack *OperandStack) {
	Gfx.PathClose()
}

func GfxLoadSprite(stack *OperandStack) {
	filename := stack.Pop().StringData
	id, err := Gfx.LoadSprite(filename)
	if err != nil {
		// How to handle errors? For now, push -1
		stack.Push(VMDataObject{Type: INTGER, IntData: -1})
		return
	}
	stack.Push(VMDataObject{Type: INTGER, IntData: int64(id)})
}

func GfxCreateSprite(stack *OperandStack) {
	heightObj := stack.Pop()
	widthObj := stack.Pop()

	var width, height int

	if widthObj.Type == REAL {
		width = int(widthObj.FloatData)
	} else {
		width = int(widthObj.IntData)
	}

	if heightObj.Type == REAL {
		height = int(heightObj.FloatData)
	} else {
		height = int(heightObj.IntData)
	}

	id := Gfx.CreateSprite(width, height)
	stack.Push(VMDataObject{Type: INTGER, IntData: int64(id)})
}

func GfxDestroySprite(stack *OperandStack) {
	idObj := stack.Pop()
	var id int
	if idObj.Type == REAL {
		id = int(idObj.FloatData)
	} else {
		id = int(idObj.IntData)
	}
	Gfx.DestroySprite(id)
}

func GfxDrawSprite(stack *OperandStack) {
	yObj := stack.Pop()
	xObj := stack.Pop()
	idObj := stack.Pop()

	var x, y, id int

	if idObj.Type == REAL {
		id = int(idObj.FloatData)
	} else {
		id = int(idObj.IntData)
	}

	if xObj.Type == REAL {
		x = int(xObj.FloatData)
	} else {
		x = int(xObj.IntData)
	}

	if yObj.Type == REAL {
		y = int(yObj.FloatData)
	} else {
		y = int(yObj.IntData)
	}

	Gfx.DrawSprite(id, x, y)
}

func GfxSetSpriteRotation(stack *OperandStack) {
	angleObj := stack.Pop()
	idObj := stack.Pop()

	var id int
	var angle float64

	if idObj.Type == REAL {
		id = int(idObj.FloatData)
	} else {
		id = int(idObj.IntData)
	}

	if angleObj.Type == REAL {
		angle = angleObj.FloatData
	} else {
		angle = float64(angleObj.IntData)
	}

	Gfx.SetSpriteRotation(id, angle)
}

func GfxSetSpriteScale(stack *OperandStack) {
	syObj := stack.Pop()
	sxObj := stack.Pop()
	idObj := stack.Pop()

	var id int
	var sx, sy float64

	if idObj.Type == REAL {
		id = int(idObj.FloatData)
	} else {
		id = int(idObj.IntData)
	}

	if sxObj.Type == REAL {
		sx = sxObj.FloatData
	} else {
		sx = float64(sxObj.IntData)
	}

	if syObj.Type == REAL {
		sy = syObj.FloatData
	} else {
		sy = float64(syObj.IntData)
	}

	Gfx.SetSpriteScale(id, sx, sy)
}

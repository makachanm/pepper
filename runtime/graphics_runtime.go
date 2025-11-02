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
		width = int(widthObj.Value.(float64))
	} else {
		width = int(widthObj.Value.(int64))
	}

	if heightObj.Type == REAL {
		height = int(heightObj.Value.(float64))
	} else {
		height = int(heightObj.Value.(int64))
	}

	Gfx.Resize(width, height)
}

func GfxGetWidth(stack *OperandStack) {
	width, _ := Gfx.GetDimensions()
	stack.Push(makeIntValueObj(int64(width)))
}

func GfxGetHeight(stack *OperandStack) {
	_, height := Gfx.GetDimensions()
	stack.Push(makeIntValueObj(int64(height)))
}

func GfxSetWindowTitle(stack *OperandStack) {
	title := stack.Pop().Value.(string)
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
		r = rObj.Value.(float64)
	} else {
		r = float64(rObj.Value.(int64)) / 255.0
	}

	if gObj.Type == REAL {
		g = gObj.Value.(float64)
	} else {
		g = float64(gObj.Value.(int64)) / 255.0
	}

	if bObj.Type == REAL {
		b = bObj.Value.(float64)
	} else {
		b = float64(bObj.Value.(int64)) / 255.0
	}

	Gfx.SetSourceRGB(r, g, b)
}

func GfxSetSourceRGBA(stack *OperandStack) {
	aObj := stack.Pop()
	bObj := stack.Pop()
	gObj := stack.Pop()
	rObj := stack.Pop()

	var r, g, b, a float64

	if rObj.Type == REAL {
		r = rObj.Value.(float64)
	} else {
		r = float64(rObj.Value.(int64)) / 255.0
	}

	if gObj.Type == REAL {
		g = gObj.Value.(float64)
	} else {
		g = float64(gObj.Value.(int64)) / 255.0
	}

	if bObj.Type == REAL {
		b = bObj.Value.(float64)
	} else {
		b = float64(bObj.Value.(int64)) / 255.0
	}

	if aObj.Type == REAL {
		a = aObj.Value.(float64)
	} else {
		a = float64(aObj.Value.(int64)) / 255.0
	}

	Gfx.SetSourceRGBA(r, g, b, a)
}

func GfxDrawRect(stack *OperandStack) {
	heightObj := stack.Pop()
	widthObj := stack.Pop()
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y, width, height int

	if xObj.Type == REAL {
		x = int(xObj.Value.(float64))
	} else {
		x = int(xObj.Value.(int64))
	}

	if yObj.Type == REAL {
		y = int(yObj.Value.(float64))
	} else {
		y = int(yObj.Value.(int64))
	}

	if widthObj.Type == REAL {
		width = int(widthObj.Value.(float64))
	} else {
		width = int(widthObj.Value.(int64))
	}

	if heightObj.Type == REAL {
		height = int(heightObj.Value.(float64))
	} else {
		height = int(heightObj.Value.(int64))
	}

	Gfx.DrawRect(x, y, width, height)
}

func GfxDrawDot(stack *OperandStack) {
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y int

	if xObj.Type == REAL {
		x = int(xObj.Value.(float64))
	} else {
		x = int(xObj.Value.(int64))
	}

	if yObj.Type == REAL {
		y = int(yObj.Value.(float64))
	} else {
		y = int(yObj.Value.(int64))
	}

	Gfx.DrawDot(x, y)
}

func GfxDrawCircle(stack *OperandStack) {
	radiusObj := stack.Pop()
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y, radius int

	if xObj.Type == REAL {
		x = int(xObj.Value.(float64))
	} else {
		x = int(xObj.Value.(int64))
	}

	if yObj.Type == REAL {
		y = int(yObj.Value.(float64))
	} else {
		y = int(yObj.Value.(int64))
	}

	if radiusObj.Type == REAL {
		radius = int(radiusObj.Value.(float64))
	} else {
		radius = int(radiusObj.Value.(int64))
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
		x1 = int(x1Obj.Value.(float64))
	} else {
		x1 = int(x1Obj.Value.(int64))
	}

	if y1Obj.Type == REAL {
		y1 = int(y1Obj.Value.(float64))
	} else {
		y1 = int(y1Obj.Value.(int64))
	}

	if x2Obj.Type == REAL {
		x2 = int(x2Obj.Value.(float64))
	} else {
		x2 = int(x2Obj.Value.(int64))
	}

	if y2Obj.Type == REAL {
		y2 = int(y2Obj.Value.(float64))
	} else {
		y2 = int(y2Obj.Value.(int64))
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
		x1 = int(x1Obj.Value.(float64))
	} else {
		x1 = int(x1Obj.Value.(int64))
	}

	if y1Obj.Type == REAL {
		y1 = int(y1Obj.Value.(float64))
	} else {
		y1 = int(y1Obj.Value.(int64))
	}

	if x2Obj.Type == REAL {
		x2 = int(x2Obj.Value.(float64))
	} else {
		x2 = int(x2Obj.Value.(int64))
	}

	if y2Obj.Type == REAL {
		y2 = int(y2Obj.Value.(float64))
	} else {
		y2 = int(y2Obj.Value.(int64))
	}

	if x3Obj.Type == REAL {
		x3 = int(x3Obj.Value.(float64))
	} else {
		x3 = int(x3Obj.Value.(int64))
	}

	if y3Obj.Type == REAL {
		y3 = int(y3Obj.Value.(float64))
	} else {
		y3 = int(y3Obj.Value.(int64))
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
		x1 = int(x1Obj.Value.(float64))
	} else {
		x1 = int(x1Obj.Value.(int64))
	}

	if y1Obj.Type == REAL {
		y1 = int(y1Obj.Value.(float64))
	} else {
		y1 = int(y1Obj.Value.(int64))
	}

	if x2Obj.Type == REAL {
		x2 = int(x2Obj.Value.(float64))
	} else {
		x2 = int(x2Obj.Value.(int64))
	}

	if y2Obj.Type == REAL {
		y2 = int(y2Obj.Value.(float64))
	} else {
		y2 = int(y2Obj.Value.(int64))
	}

	if x3Obj.Type == REAL {
		x3 = int(x3Obj.Value.(float64))
	} else {
		x3 = int(x3Obj.Value.(int64))
	}

	if y3Obj.Type == REAL {
		y3 = int(y3Obj.Value.(float64))
	} else {
		y3 = int(y3Obj.Value.(int64))
	}

	if x4Obj.Type == REAL {
		x4 = int(x4Obj.Value.(float64))
	} else {
		x4 = int(x4Obj.Value.(int64))
	}

	if y4Obj.Type == REAL {
		y4 = int(y4Obj.Value.(float64))
	} else {
		y4 = int(y4Obj.Value.(int64))
	}

	Gfx.DrawBezier(x1, y1, x2, y2, x3, y3, x4, y4)
}

func GfxDrawText(stack *OperandStack) {
	text := stack.Pop().Value.(string)
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y int

	if xObj.Type == REAL {
		x = int(xObj.Value.(float64))
	} else {
		x = int(xObj.Value.(int64))
	}

	if yObj.Type == REAL {
		y = int(yObj.Value.(float64))
	} else {
		y = int(yObj.Value.(int64))
	}

	Gfx.DrawText(x, y, text)
}

func GfxSetFontFace(stack *OperandStack) {
	fontFace := stack.Pop().Value.(string)
	Gfx.SetFontFace(fontFace)
}

func GfxSetFontSize(stack *OperandStack) {
	sizeObj := stack.Pop()
	var size float64
	if sizeObj.Type == REAL {
		size = sizeObj.Value.(float64)
	} else {
		size = float64(sizeObj.Value.(int64))
	}
	Gfx.SetFontSize(size)
}

func GfxSaveToFile(stack *OperandStack) {
	filename := stack.Pop().Value.(string)
	Gfx.SaveToFile(filename)
}

func GfxFinish() {
	Gfx.Finish()
}

func GfxSetLineWidth(stack *OperandStack) {
	widthObj := stack.Pop()
	var width float64
	if widthObj.Type == REAL {
		width = widthObj.Value.(float64)
	} else {
		width = float64(widthObj.Value.(int64))
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
		x = int(xObj.Value.(float64))
	} else {
		x = int(xObj.Value.(int64))
	}

	if yObj.Type == REAL {
		y = int(yObj.Value.(float64))
	} else {
		y = int(yObj.Value.(int64))
	}

	if widthObj.Type == REAL {
		width = int(widthObj.Value.(float64))
	} else {
		width = int(widthObj.Value.(int64))
	}

	if heightObj.Type == REAL {
		height = int(heightObj.Value.(float64))
	} else {
		height = int(heightObj.Value.(int64))
	}

	Gfx.PathRectangle(x, y, width, height)
}

func GfxPathCircle(stack *OperandStack) {
	radiusObj := stack.Pop()
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y, radius int

	if xObj.Type == REAL {
		x = int(xObj.Value.(float64))
	} else {
		x = int(xObj.Value.(int64))
	}

	if yObj.Type == REAL {
		y = int(yObj.Value.(float64))
	} else {
		y = int(yObj.Value.(int64))
	}

	if radiusObj.Type == REAL {
		radius = int(radiusObj.Value.(float64))
	} else {
		radius = int(radiusObj.Value.(int64))
	}

	Gfx.PathCircle(x, y, radius)
}

func GfxPathMoveTo(stack *OperandStack) {
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y int

	if xObj.Type == REAL {
		x = int(xObj.Value.(float64))
	} else {
		x = int(xObj.Value.(int64))
	}

	if yObj.Type == REAL {
		y = int(yObj.Value.(float64))
	} else {
		y = int(yObj.Value.(int64))
	}

	Gfx.PathMoveTo(x, y)
}

func GfxPathLineTo(stack *OperandStack) {
	yObj := stack.Pop()
	xObj := stack.Pop()

	var x, y int

	if xObj.Type == REAL {
		x = int(xObj.Value.(float64))
	} else {
		x = int(xObj.Value.(int64))
	}

	if yObj.Type == REAL {
		y = int(yObj.Value.(float64))
	} else {
		y = int(yObj.Value.(int64))
	}

	Gfx.PathLineTo(x, y)
}

func GfxPathClose(stack *OperandStack) {
	Gfx.PathClose()
}

func GfxLoadSprite(stack *OperandStack) {
	filename := stack.Pop().Value.(string)
	id, err := Gfx.LoadSprite(filename)
	if err != nil {
		// How to handle errors? For now, push -1
		stack.Push(makeIntValueObj(-1))
		return
	}
	stack.Push(makeIntValueObj(int64(id)))
}

func GfxCreateSprite(stack *OperandStack) {
	heightObj := stack.Pop()
	widthObj := stack.Pop()

	var width, height int

	if widthObj.Type == REAL {
		width = int(widthObj.Value.(float64))
	} else {
		width = int(widthObj.Value.(int64))
	}

	if heightObj.Type == REAL {
		height = int(heightObj.Value.(float64))
	} else {
		height = int(heightObj.Value.(int64))
	}

	id := Gfx.CreateSprite(width, height)
	stack.Push(makeIntValueObj(int64(id)))
}

func GfxDestroySprite(stack *OperandStack) {
	idObj := stack.Pop()
	var id int
	if idObj.Type == REAL {
		id = int(idObj.Value.(float64))
	} else {
		id = int(idObj.Value.(int64))
	}
	Gfx.DestroySprite(id)
}

func GfxDrawSprite(stack *OperandStack) {
	yObj := stack.Pop()
	xObj := stack.Pop()
	idObj := stack.Pop()

	var x, y, id int

	if idObj.Type == REAL {
		id = int(idObj.Value.(float64))
	} else {
		id = int(idObj.Value.(int64))
	}

	if xObj.Type == REAL {
		x = int(xObj.Value.(float64))
	} else {
		x = int(xObj.Value.(int64))
	}

	if yObj.Type == REAL {
		y = int(yObj.Value.(float64))
	} else {
		y = int(yObj.Value.(int64))
	}

	Gfx.DrawSprite(id, x, y)
}

func GfxSetSpriteRotation(stack *OperandStack) {
	angleObj := stack.Pop()
	idObj := stack.Pop()

	var id int
	var angle float64

	if idObj.Type == REAL {
		id = int(idObj.Value.(float64))
	} else {
		id = int(idObj.Value.(int64))
	}

	if angleObj.Type == REAL {
		angle = angleObj.Value.(float64)
	} else {
		angle = float64(angleObj.Value.(int64))
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
		id = int(idObj.Value.(float64))
	} else {
		id = int(idObj.Value.(int64))
	}

	if sxObj.Type == REAL {
		sx = sxObj.Value.(float64)
	} else {
		sx = float64(sxObj.Value.(int64))
	}

	if syObj.Type == REAL {
		sy = syObj.Value.(float64)
	} else {
		sy = float64(syObj.Value.(int64))
	}

	Gfx.SetSpriteScale(id, sx, sy)
}

func GfxSetMask(stack *OperandStack) {
	yObj := stack.Pop()
	xObj := stack.Pop()
	idObj := stack.Pop()

	var x, y, id int

	if idObj.Type == REAL {
		id = int(idObj.Value.(float64))
	} else {
		id = int(idObj.Value.(int64))
	}

	if xObj.Type == REAL {
		x = int(xObj.Value.(float64))
	} else {
		x = int(xObj.Value.(int64))
	}

	if yObj.Type == REAL {
		y = int(yObj.Value.(float64))
	} else {
		y = int(yObj.Value.(int64))
	}

	Gfx.SetMask(id, x, y)
}

func GfxResetMask(stack *OperandStack) {
	Gfx.ResetMask()
}

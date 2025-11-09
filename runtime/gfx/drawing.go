package gfx

import "github.com/ungerik/go-cairo"

func (pg *PepperGraphics) SetSourceRGB(r, g, b float64) {
	pg.Surface.SetSourceRGB(r, g, b)
}

func (pg *PepperGraphics) SetSourceRGBA(r, g, b, a float64) {
	pg.Surface.SetSourceRGBA(r, g, b, a)
}

func (pg *PepperGraphics) DrawRect(x, y, width, height int) {
	pg.Surface.Rectangle(float64(x), float64(y), float64(width), float64(height))
	pg.Surface.Fill()
}

func (pg *PepperGraphics) DrawDot(x, y int) {
	pg.Surface.Rectangle(float64(x), float64(y), 1, 1)
	pg.Surface.Fill()
}

func (pg *PepperGraphics) DrawCircle(x, y, radius int) {
	pg.Surface.Arc(float64(x), float64(y), float64(radius), 0, 2*3.141592)
	pg.Surface.Fill()
}

func (pg *PepperGraphics) DrawLine(x1, y1, x2, y2 int) {
	pg.Surface.MoveTo(float64(x1), float64(y1))
	pg.Surface.LineTo(float64(x2), float64(y2))
	pg.Surface.Stroke()
}

func (pg *PepperGraphics) DrawTriangle(x1, y1, x2, y2, x3, y3 int) {
	pg.Surface.MoveTo(float64(x1), float64(y1))
	pg.Surface.LineTo(float64(x2), float64(y2))
	pg.Surface.LineTo(float64(x3), float64(y3))
	pg.Surface.ClosePath()
	pg.Surface.Fill()
}

func (pg *PepperGraphics) DrawBezier(x1, y1, x2, y2, x3, y3, x4, y4 int) {
	pg.Surface.MoveTo(float64(x1), float64(y1))
	pg.Surface.CurveTo(float64(x2), float64(y2), float64(x3), float64(y3), float64(x4), float64(y4))
	pg.Surface.Stroke()
}

func (pg *PepperGraphics) SetFontFace(fontFace string) {
	pg.Surface.SelectFontFace(fontFace, cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
}

func (pg *PepperGraphics) SetFontSize(size float64) {
	pg.Surface.SetFontSize(size)
}

func (pg *PepperGraphics) DrawText(x, y int, text string) {
	pg.Surface.MoveTo(float64(x), float64(y))
	pg.Surface.ShowText(text)
}

func (pg *PepperGraphics) SetLineWidth(width float64) {
	pg.Surface.SetLineWidth(width)
}

func (pg *PepperGraphics) Stroke() {
	pg.Surface.Stroke()
}

func (pg *PepperGraphics) Fill() {
	pg.Surface.Fill()
}

func (pg *PepperGraphics) PathRectangle(x, y, width, height int) {
	pg.Surface.Rectangle(float64(x), float64(y), float64(width), float64(height))
}

func (pg *PepperGraphics) PathCircle(x, y, radius int) {
	pg.Surface.Arc(float64(x), float64(y), float64(radius), 0, 2*3.141592)
}

func (pg *PepperGraphics) PathMoveTo(x, y int) {
	pg.Surface.MoveTo(float64(x), float64(y))
}

func (pg *PepperGraphics) PathLineTo(x, y int) {
	pg.Surface.LineTo(float64(x), float64(y))
}

func (pg *PepperGraphics) PathClose() {
	pg.Surface.ClosePath()
}

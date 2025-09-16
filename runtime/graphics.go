package runtime

import (
	"github.com/ungerik/go-cairo"
)

type CairoGraphics struct {
	Width   int
	Height  int
	Surface *cairo.Surface
}

func NewCairoGraphics(width, height int) *CairoGraphics {
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, width, height)
	return &CairoGraphics{
		Width:   width,
		Height:  height,
		Surface: surface,
	}
}
func (cg *CairoGraphics) Resize(width, height int) {
	cg.Width = width
	cg.Height = height
	cg.Surface.Finish()
	cg.Surface = cairo.NewSurface(cairo.FORMAT_ARGB32, width, height)
}

func (cg *CairoGraphics) GetDimensions() (int, int) {
	return cg.Width, cg.Height
}
func (cg *CairoGraphics) Clear() {
	cg.Surface.SetSourceRGB(0, 0, 0)
	cg.Surface.Paint()
}

func (cg *CairoGraphics) SetSourceRGB(r, g, b float64) {
	cg.Surface.SetSourceRGB(r, g, b)
}

func (cg *CairoGraphics) DrawRect(x, y, width, height int) {
	cg.Surface.Rectangle(float64(x), float64(y), float64(width), float64(height))
	cg.Surface.Fill()
}

func (cg *CairoGraphics) DrawCircle(x, y, radius int) {
	cg.Surface.Arc(float64(x), float64(y), float64(radius), 0, 2*3.141592)
	cg.Surface.Fill()
}

func (cg *CairoGraphics) DrawLine(x1, y1, x2, y2 int) {
	cg.Surface.MoveTo(float64(x1), float64(y1))
	cg.Surface.LineTo(float64(x2), float64(y2))
	cg.Surface.Stroke()
}

func (cg *CairoGraphics) DrawTriangle(x1, y1, x2, y2, x3, y3 int) {
	cg.Surface.MoveTo(float64(x1), float64(y1))
	cg.Surface.LineTo(float64(x2), float64(y2))
	cg.Surface.LineTo(float64(x3), float64(y3))
	cg.Surface.ClosePath()
	cg.Surface.Fill()
}

func (cg *CairoGraphics) DrawBezier(x1, y1, x2, y2, x3, y3, x4, y4 int) {
	cg.Surface.MoveTo(float64(x1), float64(y1))
	cg.Surface.CurveTo(float64(x2), float64(y2), float64(x3), float64(y3), float64(x4), float64(y4))
	cg.Surface.Stroke()
}

func (cg *CairoGraphics) SetFontFace(fontFace string) {
	cg.Surface.SelectFontFace(fontFace, cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
}

func (cg *CairoGraphics) SetFontSize(size float64) {
	cg.Surface.SetFontSize(size)
}

func (cg *CairoGraphics) DrawText(x, y int, text string) {
	cg.Surface.MoveTo(float64(x), float64(y))
	cg.Surface.ShowText(text)
}

func (cg *CairoGraphics) SaveToFile(filename string) {
	cg.Surface.WriteToPNG(filename)
}

func (cg *CairoGraphics) Finish() {
	cg.Surface.Finish()
}

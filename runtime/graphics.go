package runtime

type Graphics interface {
	Resize(width, height int)
	GetDimensions() (int, int)
	Clear()
	SetSourceRGB(r, g, b float64)
	DrawRect(x, y, width, height int)
	DrawCircle(x, y, radius int)
	DrawLine(x1, y1, x2, y2 int)
	DrawTriangle(x1, y1, x2, y2, x3, y3 int)
	DrawBezier(x1, y1, x2, y2, x3, y3, x4, y4 int)
	SetFontFace(fontFace string)
	SetFontSize(size float64)
	DrawText(x, y int, text string)
	SaveToFile(filename string)
	Finish()
}

package runtime

import (
	"sync"

	"github.com/ungerik/go-cairo"
	"github.com/veandco/go-sdl2/sdl"
)

type PepperGraphics struct {
	Width        int
	Height       int
	Window       *sdl.Window
	Surface      *cairo.Surface
	wg           *sync.WaitGroup
	sprites      map[int]*Sprite
	nextSpriteID int
}

func NewGraphics(width, height int, wg *sync.WaitGroup) *PepperGraphics {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Pepper Graphics", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	sdlSurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	cairoSurface := cairo.NewSurfaceFromData(sdlSurface.Data(), cairo.FORMAT_ARGB32, width, height, int(sdlSurface.Pitch))

	pg := &PepperGraphics{
		Width:        width,
		Height:       height,
		Window:       window,
		Surface:      cairoSurface,
		wg:           wg,
		sprites:      make(map[int]*Sprite),
		nextSpriteID: 0,
	}

	// Start the event loop in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for !ShouldQuit {
			event := sdl.PollEvent()
			if event != nil {
				switch e := event.(type) {
				case *sdl.QuitEvent:
					EventQueue.Enqueue(Event{Type: EventTypeQuit})
					ShouldQuit = true // Signal VM to quit
				case *sdl.MouseMotionEvent:
					EventQueue.Enqueue(Event{
						Type: EventTypeMouseMotion,
						X:    int(e.X),
						Y:    int(e.Y),
					})
				case *sdl.MouseButtonEvent:
					var eventType EventType
					if e.State == sdl.PRESSED {
						eventType = EventTypeMouseButtonDown
					} else {
						eventType = EventTypeMouseButtonUp
					}
					EventQueue.Enqueue(Event{
						Type:   eventType,
						X:      int(e.X),
						Y:      int(e.Y),
						Button: e.Button,
					})
				case *sdl.KeyboardEvent:
					var eventType EventType
					if e.State == sdl.PRESSED {
						eventType = EventTypeKeyDown
					} else {
						eventType = EventTypeKeyUp
					}
					EventQueue.Enqueue(Event{
						Type:    eventType,
						Key:     e.Keysym.Sym,
						KeyName: sdl.GetKeyName(e.Keysym.Sym),
					})
				}
			}
			// Small delay to prevent busy-waiting
		}
		pg.Surface.Finish()
		pg.Window.Destroy()
		sdl.Quit()
	}()

	return pg
}

func (pg *PepperGraphics) Resize(width, height int) {
	pg.Window.SetSize(int32(width), int32(height))
	sdlSurface, err := pg.Window.GetSurface()
	if err != nil {
		panic(err)
	}
	cairoSurface := cairo.NewSurfaceFromData(sdlSurface.Data(), cairo.FORMAT_ARGB32, width, height, int(sdlSurface.Pitch))
	pg.Surface.Finish() // Finish with the old surface
	pg.Width = width
	pg.Height = height
	pg.Surface = cairoSurface
}

func (pg *PepperGraphics) SetWindowTitle(title string) {
	pg.Window.SetTitle(title)
}

func (pg *PepperGraphics) GetDimensions() (int, int) {
	return pg.Width, pg.Height
}

func (pg *PepperGraphics) Clear() {
	pg.Surface.SetSourceRGB(0, 0, 0)
	pg.Surface.Paint()
}

func (pg *PepperGraphics) SetSourceRGB(r, g, b float64) {
	pg.Surface.SetSourceRGB(r, g, b)
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

func (pg *PepperGraphics) SaveToFile(filename string) {
	pg.Surface.WriteToPNG(filename)
}

func (pg *PepperGraphics) Finish() {
	pg.Surface.Flush()
	pg.Window.UpdateSurface()
}

// New methods

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

// Sprite methods

func (pg *PepperGraphics) LoadSprite(filename string) (int, error) {
	surface, err := cairo.NewSurfaceFromPNG(filename)
	if err == cairo.STATUS_FILE_NOT_FOUND {
		panic(filename + " file cannot be loaded!")
	}

	id := pg.nextSpriteID
	pg.nextSpriteID++

	pg.sprites[id] = &Sprite{
		Surface:  surface,
		Visible:  true,
		Rotation: 0,
		ScaleX:   1.0,
		ScaleY:   1.0,
	}

	return id, nil
}

func (pg *PepperGraphics) CreateSprite(width, height int) int {
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, width, height)

	id := pg.nextSpriteID
	pg.nextSpriteID++

	pg.sprites[id] = &Sprite{
		Surface:  surface,
		Visible:  true,
		Rotation: 0,
		ScaleX:   1.0,
		ScaleY:   1.0,
	}

	return id
}

func (pg *PepperGraphics) DestroySprite(id int) {
	sprite, ok := pg.sprites[id]
	if !ok {
		return
	}
	sprite.Surface.Finish()
	delete(pg.sprites, id)
}

func (pg *PepperGraphics) DrawSprite(id, x, y int) {
	sprite, ok := pg.sprites[id]
	if !ok || !sprite.Visible {
		return
	}

	pg.Surface.Save()
	pg.Surface.Translate(float64(x), float64(y))
	pg.Surface.Rotate(sprite.Rotation * 3.141592 / 180)
	pg.Surface.Scale(sprite.ScaleX, sprite.ScaleY)
	pg.Surface.SetSourceSurface(sprite.Surface, 0, 0)
	pg.Surface.Paint()
	pg.Surface.Restore()
}

func (pg *PepperGraphics) SetSpriteRotation(id int, angle float64) {
	sprite, ok := pg.sprites[id]
	if !ok {
		return
	}
	sprite.Rotation = angle
}

func (pg *PepperGraphics) SetSpriteScale(id int, sx, sy float64) {
	sprite, ok := pg.sprites[id]
	if !ok {
		return
	}
	sprite.ScaleX = sx
	sprite.ScaleY = sy
}

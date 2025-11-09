package gfx

import (
	"sync"

	"github.com/ungerik/go-cairo"
	"github.com/veandco/go-sdl2/sdl"
)

var ShouldQuit bool = false

type PepperGraphics struct {
	Width        int
	Height       int
	Window       *sdl.Window
	Surface      *cairo.Surface
	wg           *sync.WaitGroup
	sprites      map[int]*Sprite
	nextSpriteID int
	isMasking    bool
	maskSpriteID int
	maskX        int
	maskY        int
}

func NewPepperGraphics(width, height int, wg *sync.WaitGroup) *PepperGraphics {
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
		isMasking:    false,
		maskSpriteID: -1,
		maskX:        0,
		maskY:        0,
	}

	// Start the event loop in a separate goroutine
	wg.Add(1)
	go pg.runEventLoop()

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

func (pg *PepperGraphics) SaveToFile(filename string) {
	pg.Surface.WriteToPNG(filename)
}

func (pg *PepperGraphics) Finish() {
	pg.Surface.Flush()
	pg.Window.UpdateSurface()
}

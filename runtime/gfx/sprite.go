package gfx

import (
	"fmt"
	"log"
	"os"

	"github.com/rustyoz/svg"
	"github.com/ungerik/go-cairo"
)

type Sprite struct {
	Surface  *cairo.Surface
	IsSVG    bool
	SVG      svg.Svg
	Visible  bool
	Rotation float64
	ScaleX   float64
	ScaleY   float64
}

func (pg *PepperGraphics) LoadSprite(filename string) (int, error) {
	if len(filename) >= 4 && filename[len(filename)-4:] == ".svg" {
		svgRaw, err := os.ReadFile(filename)
		if err != nil {
			panic(filename + " file cannot be loaded!")
		}

		svgImage, err := svg.ParseSvg(string(svgRaw), "SVG", 1)
		if err != nil {
			panic(filename + " is not a valid SVG file!")
		}

		id := pg.nextSpriteID
		pg.nextSpriteID++

		pg.sprites[id] = &Sprite{
			IsSVG:    true,
			SVG:      *svgImage,
			Visible:  true,
			Rotation: 0,
			ScaleX:   1.0,
			ScaleY:   1.0,
		}

		return id, nil
	}

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
	if sprite.Surface != nil {
		sprite.Surface.Finish()
	}
	delete(pg.sprites, id)
}

func hexToRGBA(hex string) (float64, float64, float64, float64) {
	if hex == "none" {
		return 0, 0, 0, 0
	}
	var r, g, b, a uint8
	a = 255
	if len(hex) == 7 {
		fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	} else if len(hex) == 4 {
		fmt.Sscanf(hex, "#%1x%1x%1x", &r, &g, &b)
		r *= 17
		g *= 17
		b *= 17
	}
	return float64(r) / 255, float64(g) / 255, float64(b) / 255, float64(a) / 255
}

func (pg *PepperGraphics) DrawSprite(id, x, y int) {
	sprite, ok := pg.sprites[id]
	if !ok || !sprite.Visible {
		return
	}

	if sprite.IsSVG {
		pg.Surface.Save()
		pg.Surface.Translate(float64(x), float64(y))
		pg.Surface.Rotate(sprite.Rotation * 3.141592 / 180)
		pg.Surface.Scale(sprite.ScaleX, sprite.ScaleY)

		instructions, errors := sprite.SVG.ParseDrawingInstructions()
		go func() {
			for err := range errors {
				log.Printf("SVG Engine Error: %v", err)
			}
		}()

		isPathClosded := false
		pg.Surface.NewPath()
		for instr := range instructions {
			switch instr.Kind {
			case svg.MoveInstruction:
				if isPathClosded {
					pg.Surface.NewPath()
					isPathClosded = false
				}
				pg.Surface.MoveTo(instr.M[0], instr.M[1])
			case svg.LineInstruction:
				if isPathClosded {
					pg.Surface.NewPath()
					isPathClosded = false
				}
				pg.Surface.LineTo(instr.M[0], instr.M[1])
			case svg.CurveInstruction:
				if isPathClosded {
					pg.Surface.NewPath()
					isPathClosded = false
				}
				pg.Surface.CurveTo(instr.CurvePoints.C1[0], instr.CurvePoints.C1[1], instr.CurvePoints.C2[0], instr.CurvePoints.C2[1], instr.CurvePoints.T[0], instr.CurvePoints.T[1])
			case svg.CircleInstruction:
				if isPathClosded {
					pg.Surface.NewPath()
					isPathClosded = false
				}
				pg.Surface.Arc(instr.M[0], instr.M[1], *instr.Radius, 0, 2*3.141592)
			case svg.CloseInstruction:
				pg.Surface.ClosePath()
				isPathClosded = true
			case svg.PaintInstruction:
				if isPathClosded {
					pg.Surface.NewPath()
					isPathClosded = false
				}
				if instr.Fill != nil {
					r, g, b, a := hexToRGBA(*instr.Fill)
					pg.Surface.SetSourceRGBA(r, g, b, a)
					if instr.Opacity != nil {
						pg.Surface.SetSourceRGBA(r, g, b, *instr.Opacity)
					}
					pg.Surface.FillPreserve()
				}
				if instr.Stroke != nil {
					r, g, b, a := hexToRGBA(*instr.Stroke)
					pg.Surface.SetSourceRGBA(r, g, b, a)
					if instr.Opacity != nil {
						pg.Surface.SetSourceRGBA(r, g, b, *instr.Opacity)
					}
					if instr.StrokeWidth != nil {
						pg.Surface.SetLineWidth(*instr.StrokeWidth)
					}
					pg.Surface.StrokePreserve()
				}
				pg.Surface.NewPath()
			}
		}

		pg.Surface.Restore()
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

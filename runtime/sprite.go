package runtime

import "github.com/ungerik/go-cairo"

type Sprite struct {
	Surface  *cairo.Surface
	Visible  bool
	Rotation float64
	ScaleX   float64
	ScaleY   float64
}

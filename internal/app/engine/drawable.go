package engine

import (
	"image/color"

	"github.com/faiface/pixel"
)

type Drawable struct {
	Position pixel.Vec
	Color    color.Color
}

package engine

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	win      *pixelgl.Window
	Position pixel.Vec
	Zoom     float64
	Matrix   pixel.Matrix
}

func NewCamera(win *pixelgl.Window, pos pixel.Vec) *Camera {
	return &Camera{
		win:      win,
		Position: pos,
		Zoom:     1.0,
		Matrix:   pixel.IM,
	}
}

func (c *Camera) Update() {
	screencenter := c.win.Bounds().Center()

	movepos := pixel.V(
		math.Floor(-c.Position.X),
		math.Floor(-c.Position.Y),
	).Add(screencenter)

	c.Matrix = pixel.IM.Moved(movepos).Scaled(screencenter, c.Zoom)
}

package engine

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	win       *pixelgl.Window
	Position  pixel.Vec
	MoveSpeed float64
	Zoom      float64
	Matrix    pixel.Matrix
}

func NewCamera(win *pixelgl.Window, pos pixel.Vec) *Camera {
	return &Camera{
		win:       win,
		Position:  pos,
		MoveSpeed: 3.0,
		Zoom:      1.0,
		Matrix:    pixel.IM,
	}
}

func (c *Camera) handleInput() {
	if c.win.Pressed(pixelgl.KeyA) {
		c.Position.X -= c.MoveSpeed * c.Zoom
	}
	if c.win.Pressed(pixelgl.KeyD) {
		c.Position.X += c.MoveSpeed * c.Zoom
	}
	if c.win.Pressed(pixelgl.KeyW) {
		c.Position.Y += c.MoveSpeed * c.Zoom
	}
	if c.win.Pressed(pixelgl.KeyS) {
		c.Position.Y -= c.MoveSpeed * c.Zoom
	}
}

func (c *Camera) Update() {
	c.handleInput()
	screencenter := c.win.Bounds().Center()

	movepos := pixel.V(
		math.Floor(-c.Position.X),
		math.Floor(-c.Position.Y),
	).Add(screencenter)

	c.Matrix = pixel.IM.Moved(movepos).Scaled(screencenter, c.Zoom)
}

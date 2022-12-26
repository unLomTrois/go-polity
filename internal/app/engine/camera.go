package engine

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	zoomspeed = 10.0
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
		MoveSpeed: 100,
		Zoom:      1.0,
		Matrix:    pixel.IM,
	}
}

func (c *Camera) handleInput(dt float64) {
	if c.win.Pressed(pixelgl.KeyA) {
		c.Position.X -= c.MoveSpeed * dt * c.Zoom
	}
	if c.win.Pressed(pixelgl.KeyD) {
		c.Position.X += c.MoveSpeed * dt * c.Zoom
	}
	if c.win.Pressed(pixelgl.KeyW) {
		c.Position.Y += c.MoveSpeed * dt * c.Zoom
	}
	if c.win.Pressed(pixelgl.KeyS) {
		c.Position.Y -= c.MoveSpeed * dt * c.Zoom
	}
}

func (c *Camera) handleScroll(dt float64) {
	// camera inputs
	scroll := c.win.MouseScroll().Y
	if scroll != 0 {
		c.Zoom += zoomspeed * dt * scroll
		if c.Zoom < 1 {
			c.Zoom = 1
			c.Position = c.win.Bounds().Center()
		}
	}
}

func (c *Camera) Update(dt float64, is_imgui_hovered bool) {
	c.handleInput(dt)
	if !is_imgui_hovered {
		c.handleScroll(dt)
	}
	screencenter := c.win.Bounds().Center()

	movepos := pixel.V(
		math.Floor(-c.Position.X),
		math.Floor(-c.Position.Y),
	).Add(screencenter)

	c.Matrix = pixel.IM.Moved(movepos).Scaled(screencenter, c.Zoom)
}

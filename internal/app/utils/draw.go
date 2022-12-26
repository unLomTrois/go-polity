package utils

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func DrawLine(imd *imdraw.IMDraw, u pixel.Vec, v pixel.Vec, color color.Color) {
	imd.Color = color
	imd.Push(u)
	imd.Push(v)
	imd.Line(1)
}

func DrawBounds(imd *imdraw.IMDraw, bounds pixel.Rect, color color.Color) {
	DrawRectangle(imd, bounds, color, 1)
}

func DrawRectangle(imd *imdraw.IMDraw, bounds pixel.Rect, color color.Color, thickness float64) {
	imd.Color = color
	imd.Push(bounds.Min)
	imd.Push(bounds.Max)
	imd.Rectangle(thickness)
}

func DrawSquare(imd *imdraw.IMDraw, pos pixel.Vec, size float64, color color.Color, thickness float64) {
	imd.Color = color
	imd.Push(pos.Add(pixel.Vec{X: size, Y: size}))
	imd.Push(pos.Sub(pixel.Vec{X: size, Y: size}))
	imd.Rectangle(thickness)
}

func DrawCircle(
	imd *imdraw.IMDraw,
	pos pixel.Vec,
	radius float64,
	color color.Color,
	thickness float64,
) {
	imd.Color = color
	imd.Push(pos)
	imd.Circle(radius, thickness)
}

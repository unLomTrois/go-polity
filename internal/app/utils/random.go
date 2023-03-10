package utils

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/faiface/pixel"
)

func RandBetween(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandPosition(bounds pixel.Rect) pixel.Vec {
	return pixel.V(
		RandBetween(bounds.Min.X, bounds.Max.X),
		RandBetween(bounds.Min.Y, bounds.Max.Y),
	)
}

func RandColor() color.Color {
	return color.RGBA{
		R: uint8(RandBetween(0, 255)),
		G: uint8(RandBetween(0, 255)),
		B: uint8(RandBetween(0, 255)),
		A: 255,
	}
}

func RandomNiceColor() pixel.RGBA {
	var (
		len, r, g, b float64
	)
	for len == 0 {
		r = rand.Float64()
		g = rand.Float64()
		b = rand.Float64()
		len = math.Sqrt(r*r + g*g + b*b)
	}
	return pixel.RGB(r/len, g/len, b/len)
}

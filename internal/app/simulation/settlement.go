package sim

import "github.com/faiface/pixel"

type Settlement struct {
	Name       string
	Position   pixel.Vec
	Population uint32 // max population is like 4 294 967 295
}

func NewSettlement(name string, pos pixel.Vec, population uint32) *Settlement {
	return &Settlement{
		Name:       name,
		Position:   pos,
		Population: population,
	}
}

package sim

import (
	"image/color"
	"polity/internal/app/engine"
	"polity/internal/app/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type SettlementType uint8

const (
  Village SettlementType = iota
  City
)

type Settlement struct {
	Name string
	Population uint32 // max population is like 4 294 967 295
  Type SettlementType
	engine.Drawable
}

func NewSettlement(name string, settlement_type SettlementType, pos pixel.Vec, population uint32, color color.Color) *Settlement {
	return &Settlement{
		Name:       name,
		Population: population,
		Type:       settlement_type,
		Drawable:   engine.Drawable{Position: pos, Color: color},
	}
}

func (s *Settlement) Draw(imd *imdraw.IMDraw) {
  utils.DrawSquare(imd, s.Position, 2, s.Color)

  // utils.DrawCircle(imd, s.Position, 10, s.Color, 0)
}

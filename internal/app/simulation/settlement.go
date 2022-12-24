package sim

import (
	"image/color"
	"math"
	"polity/internal/app/engine"
	"polity/internal/app/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type SettlementType uint8

const (
  MaxPopulation = 1_000_000
)

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

func calcSizeSin(x float64) float64 {
  return math.Sin((math.Pi*x)/2)*5
}

func calcSize(pop uint32, max uint32) float64 {
  relation := float64(pop) / float64(max)
  size := calcSizeSin(relation)
  return size
}

func (s *Settlement) Draw(imd *imdraw.IMDraw) {
  if s.Type == City {
    size := 4 + calcSize(s.Population, MaxPopulation)
    utils.DrawSquare(imd, s.Position, size, s.Color)
  }
  if s.Type == Village {
    size := 2 + calcSize(s.Population, MaxPopulation/2)
    utils.DrawCircle(imd, s.Position, size, s.Color, size/2)
  }
}

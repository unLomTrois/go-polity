package sim

import (
	"image/color"
	"math"
	"polity/internal/app/engine"
	"polity/internal/app/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const (
	MaxPopulation = 1_000_000
)

type SettlementType string

const (
	Tribe SettlementType = "Tribe"
	City  SettlementType = "City"
)

type Settlement struct {
	Name       string
	Population uint32 // max population is like 4 294 967 295
	Type       SettlementType
	Size       float64
	engine.Drawable
}

func NewSettlement(name string, settlement_type SettlementType, pos pixel.Vec, population uint32, color color.Color) *Settlement {
	var size float64
	if settlement_type == City {
		size = 4 + calcSize(population, MaxPopulation)
	}
	if settlement_type == Tribe {
		size = 3 + calcSize(population, MaxPopulation)
	}

	return &Settlement{
		Name:       name,
		Population: population,
		Type:       settlement_type,
		Size:       size,
		Drawable:   engine.Drawable{Position: pos, Color: color},
	}
}

func calcSizeSin(x float64) float64 {
	return math.Sin((math.Pi*x)/2) * 5
}

func calcSize(pop uint32, max uint32) float64 {
	relation := float64(pop) / float64(max)
	size := calcSizeSin(relation)
	return size
}

func (s *Settlement) Draw(imd *imdraw.IMDraw) {
	if s.Type == City {
		utils.DrawSquare(imd, s.Position, s.Size, s.Color, 0)
	}
	if s.Type == Tribe {
		utils.DrawCircle(imd, s.Position, s.Size, s.Color, s.Size/2)
	}
}

func (s *Settlement) Pos() pixel.Vec {
	return s.Position
}

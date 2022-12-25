package sim

import (
	"polity/internal/app/names"
	"polity/internal/app/utils"

	"github.com/faiface/pixel"
)

func GenerateSettlements(bounds pixel.Rect) []*Settlement {
	arr := []*Settlement{}

	// generate tribes
	for i := 0; i < 900; i++ {
		tribe := NewSettlement(
			names.GenerateName(), Tribe, utils.RandPosition(bounds),
			uint32(utils.RandBetween(30, 1_000)),
			utils.RandomNiceColor(),
		)
		arr = append(arr, tribe)
	}
	// generate cities
	for i := 0; i < 100; i++ {
		tribe := NewSettlement(
			names.GenerateName(), City, utils.RandPosition(bounds),
			uint32(utils.RandBetween(1_000, MaxPopulation/3)),
			utils.RandomNiceColor(),
		)
		arr = append(arr, tribe)
	}

	return arr
}

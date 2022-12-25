package main

import (
	"math/rand"
	"polity/internal/app/names"
	sim "polity/internal/app/simulation"
	"polity/internal/app/utils"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "polity",
		Bounds: pixel.R(0, 0, 2000, 1000),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	check(err)

	gameloop(win)
}

func gameloop(win *pixelgl.Window) {
	// cmap := cellmap.New(win.Bounds())

	imd := imdraw.New(nil)

	// bounds of simulation
	// simbounds := win.Bounds()

	arr := []*sim.Settlement{}

	// generate tribes
	for i := 0; i < 900; i++ {
		tribe := sim.NewSettlement(
			"Ur", sim.Tribe, utils.RandPosition(win.Bounds()),
			uint32(utils.RandBetween(30, 1_000)),
			utils.RandomNiceColor(),
		)
		arr = append(arr, tribe)
	}
	// generate cities
	for i := 0; i < 100; i++ {
		tribe := sim.NewSettlement(
			names.GenerateCityName(), sim.City, utils.RandPosition(win.Bounds()),
			uint32(utils.RandBetween(1_000, sim.MaxPopulation/3)),
			utils.RandomNiceColor(),
		)
		arr = append(arr, tribe)
	}

	for !win.Closed() {

		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()

		for _, s := range arr {
			s.Draw(imd)
		}

		imd.Draw(win)

		win.Update()
	}
}

package main

import (
	"math/rand"
	sim "polity/internal/app/simulation"
	"polity/internal/app/utils"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	pixelgl.Run(run)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "polity",
		Bounds: pixel.R(0, 0, 2000, 1000),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	check(err)

	// cmap := cellmap.New(win.Bounds())

	imd := imdraw.New(nil)

	// bounds of simulation
	// simbounds := win.Bounds()

  arr := []*sim.Settlement{}

  for i := 0; i < 100; i++ {
    settlement := sim.NewSettlement(
      "Ur", sim.City, utils.RandPosition(win.Bounds()),
      uint32(utils.RandBetween(100, sim.MaxPopulation/3)),
      utils.RandColor(),
    )

    if settlement.Population < sim.MaxPopulation / 5 && rand.Float32() > 0.2 {
      settlement.Type = sim.Village
    }


    arr = append(arr, settlement)
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

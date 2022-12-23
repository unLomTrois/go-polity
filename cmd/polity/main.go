package main

import (
	"fmt"
	"math"
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
		Bounds: pixel.R(0, 0, 1024, 720),
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
    x := float64(i)/100
    y := math.Exp(x) / math.Pow((math.Exp(x) + 1), 2.0)
    fmt.Println(i, x, y)

    settlement := sim.NewSettlement(
      "Ur", sim.City, pixel.Vec{
      	X: x * win.Bounds().W(),
      	Y: y * win.Bounds().H(),
      },
      1_000,
      utils.RandColor(),
    )
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

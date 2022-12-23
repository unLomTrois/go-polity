package main

import (
	"math/rand"
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

	for !win.Closed() {

		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()

		win.Update()
	}
}

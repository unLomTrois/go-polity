package main

import (
	"log"
	"math/rand"
	"polity/internal/app/engine"
	sim "polity/internal/app/simulation"
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
	imd := imdraw.New(nil)

	arr := sim.GenerateSettlements(win.Bounds())

	camera := engine.NewCamera(win, win.Bounds().Center())

	zoomspeed := 0.2

	for !win.Closed() {
		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()

		// camera inputs
		scroll := win.MouseScroll().Y
		if scroll != 0 {
			camera.Zoom += zoomspeed * scroll
			if camera.Zoom < 1 {
				camera.Zoom = 1
				camera.Position = win.Bounds().Center()
			}
			log.Println(camera.Zoom)
		}

		// cam
		camera.Update()
		win.SetMatrix(camera.Matrix)

		// drawing
		for _, s := range arr {
			s.Draw(imd)
		}

		imd.Draw(win)

		win.SetMatrix(pixel.IM)
		win.Update()
	}
}

package main

import (
	"math/rand"
	"polity/internal/app/engine"
	"polity/internal/app/sim"
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
	imd := imdraw.New(nil)

	arr := sim.GenerateSettlements(win.Bounds())

	camera := engine.NewCamera(win, win.Bounds().Center())

	for !win.Closed() {
		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()

		// cam
		camera.Update()
		win.SetMatrix(camera.Matrix)

		// drawing
		for _, s := range arr {
			s.Draw(imd)
		}

		utils.DrawBounds(imd, win.Bounds(), colornames.White)
		imd.Draw(win)

		win.SetMatrix(pixel.IM)
		win.Update()
	}
}

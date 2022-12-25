package main

import (
	"log"
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

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()

		// cam
		camera.Update(dt)
		win.SetMatrix(camera.Matrix)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			log.Println(camera.Matrix.Unproject(win.MousePosition()))
		}
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

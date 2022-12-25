package main

import (
	"fmt"
	"math/rand"
	"polity/internal/app/engine"
	"polity/internal/app/sim"
	"polity/internal/app/utils"
	"time"

	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/inkyblackness/imgui-go"
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

	ui := pixelui.NewUI(win, 0)
	defer ui.Destroy()
	// ui.AddTTFFont("03b04.ttf", 16)
	imgui.CurrentIO().SetFontGlobalScale(2)
	camera := engine.NewCamera(win, win.Bounds().Center())

	last := time.Now()
	for !win.Closed() {
		ui.NewFrame()
		dt := time.Since(last).Seconds()
		last = time.Now()

		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()
		imgui.ShowDemoWindow(nil)

		imgui.Begin("Image Test")

		imgui.Text(fmt.Sprintf("%.2f", dt))
		imgui.End()

		// cam
		camera.Update(dt)
		win.SetMatrix(camera.Matrix)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			// log.Println()
			// log.Println(camera.Matrix.Unproject(win.MousePosition()))
		}
		// drawing
		for _, s := range arr {
			s.Draw(imd)
		}

		utils.DrawBounds(imd, win.Bounds(), colornames.White)
		imd.Draw(win)

		ui.Draw(win)

		win.SetMatrix(pixel.IM)
		win.Update()
	}
}

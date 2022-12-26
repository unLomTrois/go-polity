package main

import (
	"fmt"
	"log"
	"math/rand"
	"polity/internal/app/engine"
	"polity/internal/app/quadtree"
	"polity/internal/app/sim"
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
		Bounds: pixel.R(0, 0, 1600, 900),
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

	list := []string{"Ur", "Uruk", "Nom", "Kiev", "Moscow"}

	qt := quadtree.NewQuadTree2(win.Bounds())
	for index := range arr {
		qt.Insert(arr[index])
	}

	var kek int32 = 0
	var is_quadtree_visible = false

	var selected_settlement *sim.Settlement = nil

	last := time.Now()
	for !win.Closed() {
		ui.NewFrame()
		dt := time.Since(last).Seconds()
		last = time.Now()

		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()
		imgui.ShowDemoWindow(nil)

		imgui.Begin("Main Settings")
		imgui.Text(fmt.Sprintf("%.2f", dt))
		imgui.Checkbox("show quadtree", &is_quadtree_visible)
		imgui.ListBox("cities", &kek, list)
		imgui.End()

		// cam
		camera.Update(dt)
		win.SetMatrix(camera.Matrix)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			// log.Println()
			mousepos := camera.Matrix.Unproject(win.MousePosition())
			query := qt.Query(pixel.R(mousepos.X-5, mousepos.Y-5, mousepos.X+5, mousepos.Y+5))
			if len(query) > 0 {
				selected_settlement = query[0]
			}
			log.Println(selected_settlement)
		}

		if selected_settlement != nil {
			imgui.Begin("Settlement Details")
			imgui.Text("Name: " + selected_settlement.Name)
			imgui.Text("Type: " + string(selected_settlement.Type))
			imgui.Text(fmt.Sprintf("Population: %d", selected_settlement.Population))
			imgui.End()
		}

		// drawing
		for _, s := range arr {
			s.Draw(imd)
		}

		// utils.DrawBounds(imd, win.Bounds(), colornames.White)
		if is_quadtree_visible {
			qt.Show(imd, colornames.White)
		}

		imd.Draw(win)

		ui.Draw(win)

		win.SetMatrix(pixel.IM)
		win.Update()
	}
}

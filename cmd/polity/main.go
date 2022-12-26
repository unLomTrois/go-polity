package main

import (
	"fmt"
	"math/rand"
	"polity/internal/app/engine"
	"polity/internal/app/quadtree"
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
		Bounds: pixel.R(0, 0, 1600, 900),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	check(err)

	gameloop(win)
}

type Settings struct {
	is_quadtree_visible bool
}

func NewSettings() *Settings {
	return &Settings{
		is_quadtree_visible: false,
	}
}

func gameloop(win *pixelgl.Window) {
	imd := imdraw.New(nil)

	ui := pixelui.NewUI(win, 0)
	defer ui.Destroy()
	// ui.AddTTFFont("03b04.ttf", 16)
	imgui.CurrentIO().SetFontGlobalScale(2)
	camera := engine.NewCamera(win, win.Bounds().Center())

	settlements := sim.GenerateSettlements(win.Bounds())
	qt := quadtree.NewQuadTree2(win.Bounds())
	for _, s := range settlements {
		qt.Insert(s)
	}

	settings := NewSettings()

	var selected_settlement *sim.Settlement = nil
	is_imgui_hovered := false
	last := time.Now()
	for !win.Closed() {
		ui.NewFrame()
		is_imgui_hovered = imgui.CurrentIO().WantCaptureMouse()
		dt := time.Since(last).Seconds()
		last = time.Now()

		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()

		// cam
		camera.Update(dt, is_imgui_hovered)
		win.SetMatrix(camera.Matrix)

		// ui
		// imgui.ShowDemoWindow(nil)
		ShowSettingsWindow(dt, settings)
		ShowSettlementDetailsWindow(selected_settlement)

		if !is_imgui_hovered && win.JustPressed(pixelgl.MouseButtonLeft) {
			mousepos := camera.Matrix.Unproject(win.MousePosition())
			mouse_boundary := pixel.R(mousepos.X-10, mousepos.Y-10, mousepos.X+10, mousepos.Y+10)
			query := qt.Query(mouse_boundary)
			if len(query) == 1 {
				selected_settlement = query[0]
			}
			// find the nearest
			if len(query) > 1 {
				selected_settlement = query[0]
				for _, settlement := range query {
					if mousepos.To(settlement.Position).Len() < mousepos.To(selected_settlement.Position).Len() {
						selected_settlement = settlement
					}
				}
			}
		}

		// drawing
		DrawSettlements(settlements, imd, selected_settlement)

		// utils.DrawBounds(imd, win.Bounds(), colornames.White)
		if settings.is_quadtree_visible {
			qt.Show(imd, colornames.White)
		}

		imd.Draw(win)

		ui.Draw(win)

		win.SetMatrix(pixel.IM)
		win.Update()
	}
}

func DrawSettlements(settlements []*sim.Settlement, imd *imdraw.IMDraw, selected_settlement *sim.Settlement) {
	for _, s := range settlements {
		s.Draw(imd)
		if selected_settlement == s {
			if s.Type == sim.City {
				utils.DrawSquare(imd, s.Position, s.Size+1, colornames.Red, 1)
			}
			if s.Type == sim.Tribe {
				utils.DrawCircle(imd, s.Pos(), s.Size+2, colornames.Red, 1)
			}
		}
	}
}

func ShowSettingsWindow(dt float64, settings *Settings) {
	imgui.SetNextWindowPos(imgui.Vec2{
		X: 10,
		Y: 10,
	})
	imgui.BeginV("Settings", nil, imgui.WindowFlagsAlwaysAutoResize)
	imgui.Text(fmt.Sprintf("%.2f", dt))
	imgui.Checkbox("show quadtree", &settings.is_quadtree_visible)
	imgui.End()
}

func ShowSettlementDetailsWindow(selected_settlement *sim.Settlement) {
	if selected_settlement != nil {
		imgui.Begin("Settlement Details")
		imgui.Text("Name: " + selected_settlement.Name)
		imgui.Text("Type: " + string(selected_settlement.Type))
		imgui.Text(fmt.Sprintf("Population: %d", selected_settlement.Population))
		imgui.End()
	}
}

package main

import (
	"math/rand"
	"polity/internal/app/engine"
	"polity/internal/app/gui"
	"polity/internal/app/quadtree"
	"polity/internal/app/sim"
	"polity/internal/app/utils"
	"time"

	"github.com/dusk125/pixelui"
	"github.com/dusk125/pixelutils"
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

	ui := pixelui.NewUI(win, 0)
	defer ui.Destroy()
	imgui.CurrentIO().SetFontGlobalScale(2)
	camera := engine.NewCamera(win, win.Bounds().Center())

	settlements := sim.GenerateSettlements(win.Bounds())
	qt := quadtree.NewQuadTree2(win.Bounds())
	for _, s := range settlements {
		qt.Insert(s)
	}

	settings := engine.NewSettings()

	var selected_settlement *sim.Settlement = nil
	var is_settlement_selected bool = selected_settlement != nil

	is_imgui_hovered := false
	ticker := pixelutils.NewTicker(120)
	for !win.Closed() {
		ui.NewFrame()
		is_imgui_hovered = imgui.CurrentIO().WantCaptureMouse()
		dt, fps := ticker.Tick()

		// отрисовка
		win.Clear(colornames.Black)
		imd.Clear()

		// cam
		camera.Update(dt, is_imgui_hovered)
		win.SetMatrix(camera.Matrix)

		// ui
		imgui.ShowDemoWindow(nil)
		gui.ShowDebugWindow(fps, settings)
		if is_settlement_selected {
			gui.ShowSettlementDetailsWindow(&is_settlement_selected, selected_settlement)
		}

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
			is_settlement_selected = selected_settlement != nil
		}

		// drawing
		DrawSettlements(settlements, imd, selected_settlement)

		// utils.DrawBounds(imd, win.Bounds(), colornames.White)
		if settings.Is_quadtree_visible {
			qt.Show(imd, colornames.White)
		}

		imd.Draw(win)

		ui.Draw(win)

		win.SetMatrix(pixel.IM)
		win.Update()
		ticker.Wait()
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

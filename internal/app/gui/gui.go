package gui

import (
	"fmt"
	"polity/internal/app/engine"
	"polity/internal/app/sim"

	"github.com/inkyblackness/imgui-go"
)

func ShowDebugWindow(fps float64, settings *engine.Settings) {
	imgui.SetNextWindowPos(imgui.Vec2{
		X: 10,
		Y: 10,
	})
	imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{
		X: 10,
		Y: 10,
	})
	imgui.BeginV("Debug", nil, imgui.WindowFlagsAlwaysAutoResize)

	imgui.Text(fmt.Sprintf("FPS: %.2f", fps))
	imgui.Checkbox("show quadtree", &settings.Is_quadtree_visible)
	imgui.PopStyleVar()
	imgui.End()
}

func ShowSettlementDetailsWindow(is_open *bool, selected_settlement *sim.Settlement) {
	if !imgui.BeginV("Settlement Details", is_open, 0) {
		imgui.End()
	} else {
		imgui.Text("Name: " + selected_settlement.Name)
		imgui.Text("Type: " + string(selected_settlement.Type))
		imgui.Text(fmt.Sprintf("Population: %d", selected_settlement.Population))
		imgui.End()
	}
}

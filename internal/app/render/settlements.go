package render

import (
	"polity/internal/app/sim"
	"polity/internal/app/utils"

	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

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

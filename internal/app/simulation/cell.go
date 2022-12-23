package sim

import (
	"image/color"
	"math"
	"polity/internal/app/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Cell struct {
	Position  pixel.Vec
	Color     color.Color
	Radius    float64
	Direction float64
	Speed     float64
}

func NewCell(position pixel.Vec, color color.Color, radius float64) *Cell {
	return &Cell{
		position,
		color,
		radius,
		utils.RandBetween(-2*math.Pi, 2*math.Pi),
		utils.RandBetween(0, 1),
	}
	// utils.RandBetween(-2*math.Pi, 2*math.Pi)
}

func GenerateCells(count int, bounds pixel.Rect) []*Cell {
	var cells []*Cell
	for i := 0; i < count; i++ {
		cells = append(
			cells,
			NewCell(utils.RandPosition(bounds), utils.RandColor(), utils.RandBetween(1, 3)),
		)
	}
	return cells
}

func (c *Cell) NextPosition() pixel.Vec {
	unitVec := pixel.Unit(c.Direction).Scaled(c.Speed)

	// fmt.Println("unitvec", unitVec)

	return pixel.V(c.Position.X+unitVec.X, c.Position.Y+unitVec.Y)
}

func (c *Cell) Move() {
	nextpos := c.NextPosition()
	// fmt.Println(c.Direction, math.Sin(c.Position.X), math.Cos(c.Position.Y))
	// fmt.Println(nextpos)

	c.Position.X = nextpos.X
	c.Position.Y = nextpos.Y

	// c.Direction -= math.Sin(c.Position.X) + math.Cos(c.Direction)
	c.Direction += utils.RandBetween(-0.01, 0.01)

	if c.Direction <= -math.Pi || c.Direction >= math.Pi {
		c.Direction += utils.RandBetween(-0.01, 0.01)
	}
}

// Перемещает клетку через границу карты на другую сторону
func (c *Cell) CrossBorder(bounds pixel.Rect) {
	if !bounds.Contains(c.Position) {
		lin := pixel.L(c.Position, c.NextPosition())

		intersec := bounds.IntersectionPoints(lin.Scaled(5))

		if len(intersec) > 0 {
			for i, e := range bounds.Edges() {
				// fmt.Println(i, e)
				if _, ok := lin.Scaled(5).Intersect(e); ok {
					if i == 0 {
						c.Position = pixel.V(bounds.Max.X, intersec[0].Y)
					}
					if i == 2 {
						c.Position = pixel.V(bounds.Min.X, intersec[0].Y)
					}
					if i == 1 {
						c.Position = pixel.V(intersec[0].X, bounds.Min.Y)
					}
					if i == 3 {
						c.Position = pixel.V(intersec[0].X, bounds.Max.Y)
					}
				}
			}
		}
	}
}

// отрисовывает клетку, должно идти до imd.Draw(win)
func (c *Cell) Draw(imd *imdraw.IMDraw) {
	utils.DrawCircle(imd, c.Position, c.Radius, c.Color, 0)
}

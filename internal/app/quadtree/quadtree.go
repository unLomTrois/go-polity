package quadtree

import (
	"image/color"
	"polity/internal/app/sim"
	"polity/internal/app/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type QuadTree struct {
	is_divided bool
	capacity   int
	points     []*sim.Cell
	boundary   pixel.Rect
	// children
	nw *QuadTree
	ne *QuadTree
	sw *QuadTree
	se *QuadTree
}

func NewQuadTree(boundary pixel.Rect) *QuadTree {
	return &QuadTree{
		is_divided: false,
		capacity:   4,
		points:     make([]*sim.Cell, 0),
		boundary:   boundary,
		nw:         nil,
		ne:         nil,
		sw:         nil,
		se:         nil,
	}
}

func (qt *QuadTree) InsertMap(cells map[**sim.Cell]*sim.Cell) {
	for _, c := range cells {
		qt.Insert(c)
	}
}

func (qt *QuadTree) Insert(cell *sim.Cell) bool {
	if !qt.boundary.Contains(cell.Position) {
		return false
	}

	if !qt.is_divided {
		if len(qt.points) < qt.capacity {
			qt.points = append(qt.points, cell)
			// fmt.Println("insert point", cell)

			if len(qt.points) == qt.capacity {
				qt.Subdivide()
			}

			return true
		}
	}
	if qt.nw.Insert(cell) {
		return true
	}
	if qt.ne.Insert(cell) {
		return true
	}
	if qt.sw.Insert(cell) {
		return true
	}
	if qt.se.Insert(cell) {
		return true
	}

	// fmt.Println("try to insert point into children")
	return false
}

func (qt *QuadTree) Subdivide() bool {
	qt.is_divided = true

	// переписать
	qt.nw = NewQuadTree(
		pixel.R(
			qt.boundary.Center().X-qt.boundary.W()/2,
			qt.boundary.Center().Y,
			qt.boundary.Center().X,
			qt.boundary.Max.Y,
		),
	)
	qt.ne = NewQuadTree(
		pixel.R(
			qt.boundary.Center().X,
			qt.boundary.Center().Y,
			qt.boundary.Max.X,
			qt.boundary.Max.Y,
		),
	)
	qt.sw = NewQuadTree(
		pixel.R(
			qt.boundary.Center().X-qt.boundary.W()/2,
			qt.boundary.Min.Y,
			qt.boundary.Center().X,
			qt.boundary.Center().Y,
		),
	)
	qt.se = NewQuadTree(
		pixel.R(
			qt.boundary.Center().X,
			qt.boundary.Min.Y,
			qt.boundary.Center().X+qt.boundary.W()/2,
			qt.boundary.Center().Y,
		),
	)

	ret := false
	for _, p := range qt.points {
		ret = qt.Insert(p)
	}

	qt.points = nil

	return ret
}

func (qt *QuadTree) Show(imd *imdraw.IMDraw, color color.Color) {
	utils.DrawBounds(imd, qt.boundary, color)

	if qt.is_divided {
		qt.nw.Show(imd, colornames.Red)
		qt.ne.Show(imd, colornames.Blue)
		qt.sw.Show(imd, colornames.Green)
		qt.se.Show(imd, colornames.Yellow)
	}
}

func (qt *QuadTree) Query(boundary pixel.Rect) (cells []*sim.Cell) {
	// cells := make([]*sim.Cell, 0)

	if !qt.boundary.Intersects(boundary) {
		return
	}

	for _, p := range qt.points {
		if boundary.Contains(p.Position) {
			// fmt.Println(p)
			cells = append(cells, p)
		}
	}

	if !qt.is_divided {
		return
	}

	cells = append(cells, qt.nw.Query(boundary)...)
	cells = append(cells, qt.ne.Query(boundary)...)
	cells = append(cells, qt.sw.Query(boundary)...)
	cells = append(cells, qt.se.Query(boundary)...)

	return
}

func (qt *QuadTree) Update(cells map[**sim.Cell]*sim.Cell) {
	qt.clear()
	qt.InsertMap(cells)
}

func (qt *QuadTree) clear() {
	qt.nw = nil
	qt.ne = nil
	qt.sw = nil
	qt.se = nil
	qt.is_divided = false
	qt.points = make([]*sim.Cell, 0)
}

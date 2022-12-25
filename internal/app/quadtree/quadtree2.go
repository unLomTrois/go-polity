package quadtree

import (
	"image/color"
	"polity/internal/app/engine"
	"polity/internal/app/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type QuadTree2 struct {
	is_divided bool
	capacity   int
	points     []*engine.Drawable
	boundary   pixel.Rect
	// children
	nw *QuadTree2
	ne *QuadTree2
	sw *QuadTree2
	se *QuadTree2
}

func NewQuadTree2(boundary pixel.Rect) *QuadTree2 {
	return &QuadTree2{
		is_divided: false,
		capacity:   4,
		points:     make([]*engine.Drawable, 0),
		boundary:   boundary,
		nw:         nil,
		ne:         nil,
		sw:         nil,
		se:         nil,
	}
}

// func (qt *QuadTree2) InsertMap(cells map[**sim.Cell]*sim.Cell) {
// 	for _, c := range cells {
// 		qt.Insert(c)
// 	}
// }

func (qt *QuadTree2) Insert(point *engine.Drawable) bool {
	if !qt.boundary.Contains(point.Position) {
		return false
	}

	if !qt.is_divided {
		if len(qt.points) < qt.capacity {
			qt.points = append(qt.points, point)
			// fmt.Println("insert point", cell)

			if len(qt.points) == qt.capacity {
				qt.Subdivide()
			}

			return true
		}
	}
	if qt.nw.Insert(point) {
		return true
	}
	if qt.ne.Insert(point) {
		return true
	}
	if qt.sw.Insert(point) {
		return true
	}
	if qt.se.Insert(point) {
		return true
	}

	// fmt.Println("try to insert point into children")
	return false
}

func (qt *QuadTree2) Subdivide() bool {
	qt.is_divided = true

	// переписать
	qt.nw = NewQuadTree2(
		pixel.R(
			qt.boundary.Center().X-qt.boundary.W()/2,
			qt.boundary.Center().Y,
			qt.boundary.Center().X,
			qt.boundary.Max.Y,
		),
	)
	qt.ne = NewQuadTree2(
		pixel.R(
			qt.boundary.Center().X,
			qt.boundary.Center().Y,
			qt.boundary.Max.X,
			qt.boundary.Max.Y,
		),
	)
	qt.sw = NewQuadTree2(
		pixel.R(
			qt.boundary.Center().X-qt.boundary.W()/2,
			qt.boundary.Min.Y,
			qt.boundary.Center().X,
			qt.boundary.Center().Y,
		),
	)
	qt.se = NewQuadTree2(
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

func (qt *QuadTree2) Show(imd *imdraw.IMDraw, color color.Color) {
	utils.DrawBounds(imd, qt.boundary, color)

	if qt.is_divided {
		qt.nw.Show(imd, colornames.Red)
		qt.ne.Show(imd, colornames.Blue)
		qt.sw.Show(imd, colornames.Green)
		qt.se.Show(imd, colornames.Yellow)
	}
}

// func (qt *QuadTree2) Query(boundary pixel.Rect) (cells []*sim.Cell) {
// 	// cells := make([]*sim.Cell, 0)

// 	if !qt.boundary.Intersects(boundary) {
// 		return
// 	}

// 	for _, p := range qt.points {
// 		if boundary.Contains(p.Position) {
// 			// fmt.Println(p)
// 			cells = append(cells, p)
// 		}
// 	}

// 	if !qt.is_divided {
// 		return
// 	}

// 	cells = append(cells, qt.nw.Query(boundary)...)
// 	cells = append(cells, qt.ne.Query(boundary)...)
// 	cells = append(cells, qt.sw.Query(boundary)...)
// 	cells = append(cells, qt.se.Query(boundary)...)

// 	return
// }

// func (qt *QuadTree2) Update(cells map[**sim.Cell]*sim.Cell) {
// 	qt.clear()
// 	qt.InsertMap(cells)
// }

func (qt *QuadTree2) clear() {
	qt.nw = nil
	qt.ne = nil
	qt.sw = nil
	qt.se = nil
	qt.is_divided = false
	qt.points = nil
}

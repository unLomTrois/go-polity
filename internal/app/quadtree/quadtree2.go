package quadtree

import (
	"image/color"
	"polity/internal/app/sim"
	"polity/internal/app/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Placeble interface {
	Pos() pixel.Vec
}

type QuadTree2[P Placeble] struct {
	is_divided bool
	capacity   int
	points     []*P
	boundary   pixel.Rect
	// children
	nw *QuadTree2[P]
	ne *QuadTree2[P]
	sw *QuadTree2[P]
	se *QuadTree2[P]
}

func NewQuadTree2(boundary pixel.Rect) *QuadTree2[*sim.Settlement] {
	return &QuadTree2[*sim.Settlement]{
		is_divided: false,
		capacity:   4,
		points:     []**sim.Settlement{},
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

func (qt *QuadTree2[P]) Insert(point *P) bool {
	if !qt.boundary.Contains((*point).Pos()) {
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

func (qt *QuadTree2[P]) newQT(boundary pixel.Rect) *QuadTree2[P] {
	return &QuadTree2[P]{
		is_divided: false,
		capacity:   4,
		points:     nil,
		boundary:   boundary,
		nw:         nil,
		ne:         nil,
		sw:         nil,
		se:         nil,
	}
}

func (qt *QuadTree2[P]) Subdivide() bool {
	qt.is_divided = true

	// переписать
	qt.nw = qt.newQT(
		pixel.R(
			qt.boundary.Center().X-qt.boundary.W()/2,
			qt.boundary.Center().Y,
			qt.boundary.Center().X,
			qt.boundary.Max.Y,
		),
	)
	qt.ne = qt.newQT(
		pixel.R(
			qt.boundary.Center().X,
			qt.boundary.Center().Y,
			qt.boundary.Max.X,
			qt.boundary.Max.Y,
		),
	)
	qt.sw = qt.newQT(
		pixel.R(
			qt.boundary.Center().X-qt.boundary.W()/2,
			qt.boundary.Min.Y,
			qt.boundary.Center().X,
			qt.boundary.Center().Y,
		),
	)
	qt.se = qt.newQT(
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

func (qt *QuadTree2[P]) Show(imd *imdraw.IMDraw, color color.Color) {
	utils.DrawBounds(imd, qt.boundary, color)

	if qt.is_divided {
		qt.nw.Show(imd, colornames.Red)
		qt.ne.Show(imd, colornames.Blue)
		qt.sw.Show(imd, colornames.Green)
		qt.se.Show(imd, colornames.Yellow)
	}
}

func (qt *QuadTree2[P]) Query(boundary pixel.Rect) (points []*P) {
	// cells := make([]*sim.Cell, 0)

	if !qt.boundary.Intersects(boundary) {
		return nil
	}

	for _, p := range qt.points {
		if boundary.Contains((*p).Pos()) {
			points = append(points, p)
		}
	}

	if !qt.is_divided {
		return points
	}

	points = append(points, qt.nw.Query(boundary)...)
	points = append(points, qt.ne.Query(boundary)...)
	points = append(points, qt.sw.Query(boundary)...)
	points = append(points, qt.se.Query(boundary)...)

	return points
}

// func (qt *QuadTree2) Update(cells map[**sim.Cell]*sim.Cell) {
// 	qt.clear()
// 	qt.InsertMap(cells)
// }

func (qt *QuadTree2[P]) clear() {
	qt.nw = nil
	qt.ne = nil
	qt.sw = nil
	qt.se = nil
	qt.is_divided = false
	qt.points = nil
}

package engine

type Settings struct {
	Is_quadtree_visible bool
}

func NewSettings() *Settings {
	return &Settings{
		Is_quadtree_visible: false,
	}
}

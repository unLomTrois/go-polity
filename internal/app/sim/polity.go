package sim

import "image/color"

// something like country, state, etc.
type Polity struct {
	Name        string
	Settlements []*Settlement
	Color       color.Color
}

func NewPolity(name string, color color.Color, settlements []*Settlement) *Polity {
	return &Polity{
		Name:        name,
		Color:       color,
		Settlements: []*Settlement{},
	}
}

func (p *Polity) AddSettlement(s *Settlement) {
	p.Settlements = append(p.Settlements, s)
}

func (p *Polity) SetSettlements(arr []*Settlement) {
	p.Settlements = arr
}

func (p *Polity) Update() {
	for _, s := range p.Settlements {
		s.Update()
	}
}

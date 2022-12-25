/*
это структура данных для хранения клеток (sim.Cell) в виде мап указателей
ключом к мапе являются указатели к sim.Cell
для чего? чтобы удалять элемент из мапы по указателю на элемент
*/
package cellmap

import (
	"polity/internal/app/sim"

	"github.com/faiface/pixel"
)

type CellMap struct {
	m map[**sim.Cell]*sim.Cell
}

func New(bounds pixel.Rect) (cellmap *CellMap) {
	cellmap = &CellMap{m: make(map[**sim.Cell]*sim.Cell)}
	cells := sim.GenerateCells(200, bounds)

	cellmap.PutSlice(cells)

	return cellmap
}

func (m *CellMap) Put(value *sim.Cell) {
	m.m[&value] = value
}

func (m *CellMap) PutSlice(cells []*sim.Cell) {
	for _, c := range cells {
		m.Put(c)
	}
}

func (m *CellMap) Values() []*sim.Cell {
	values := make([]*sim.Cell, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

func (m *CellMap) Keys() []**sim.Cell {
	keys := make([]**sim.Cell, m.Size())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

func (m *CellMap) Size() int {
	return len(m.m)
}

func (m *CellMap) Get(key **sim.Cell) (value *sim.Cell, found bool) {
	value, found = m.m[key]
	return
}

func (m *CellMap) GetM() map[**sim.Cell]*sim.Cell {
	return m.m
}

func (m *CellMap) Remove(key **sim.Cell) {
	if !m.IsEmpty() {
		delete(m.m, key)
	}
}

func (m *CellMap) IsEmpty() bool {
	return m.Size() == 0
}

func (m *CellMap) Clear() {
	m.m = make(map[**sim.Cell]*sim.Cell)
}

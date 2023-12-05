package grid

import (
	"fmt"
)

// Position represents a cell position in the grid
type Position struct {
	Row    int
	Column int
}

// FirstColumn returns true if the position is of the first column in a given row
func (p *Position) FirstColumn() bool {
	return p.Column == 0
}

// Region represents an area in a grid
type Region struct {
	Start Position
	End   Position
}

// Size represents the size of the grid
type Size struct {
	Rows    int
	Columns int
}

// Grid wrappers a two-dimensional  slice of any type and provides common helper functions
// to manipulate grids
type Grid[T any] struct {
	Values [][]T
}

// Includes indicates if a Position is within the cell position in the grid
func (g *Grid[T]) Includes(p Position) bool {
	return p.Row >= 0 && p.Row < len(g.Values) && p.Column >= 0 && p.Column < len(g.Values[0])
}

// Value returns the cell value of a specific position
// will panic (index out of range) if the position is out of bounds
func (g *Grid[T]) Value(p Position) T {
	return g.Values[p.Row][p.Column]
}

// Cell returns the cell at a specific Position
// will panic (index out of range) if the position is out of bounds
func (g *Grid[T]) Cell(p Position) Cell[T] {
	return Cell[T]{
		Position: p,
		Value:    g.Value(p),
	}
}

func (g *Grid[T]) AppendRow(values []T) {
	if len(g.Values) != 0 {
		if size := len(g.Values[0]); size != len(values) {
			panic(fmt.Sprintf(
				"cannot add row with values %+v\nnew row must be of the same size as previous rows, got %d expected %d",
				values, len(values), size,
			))
		}
	}

	g.Values = append(g.Values, values)
}

func (g *Grid[T]) Size() Size {
	rows := len(g.Values)
	if rows == 0 {
		return Size{}
	}

	return Size{
		Rows:    rows,
		Columns: len(g.Values[0]),
	}
}

// Rows Returns the underlying two-dimensional array
// In many cases Iterator might be a better option
func (g *Grid[T]) Rows() [][]T {
	return g.Values
}

// Iterator returns a grid.Iterator that provides a standard way to iterate over the grid cells
func (g *Grid[T]) Iterator() Iterator[T] {
	return Iterator[T]{g: g}
}

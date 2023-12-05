package grid

import (
	"fmt"
	"testing"
)

func ExampleGrid_Iterator() {
	g := Grid[int]{Values: [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}

	for i := g.Iterator(); i.Next(); {
		fmt.Println(i.Value())
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}

func ExampleGrid_Rows() {
	g := Grid[int]{Values: [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}

	// In many cases g.Iterator would fit better
	fmt.Println(g.Rows())

	// Output: [[1 2 3] [4 5 6] [7 8 9]]
}

func ExampleGrid_Includes() {
	g := Grid[int]{Values: [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}

	fmt.Println(g.Includes(Position{Row: 0, Column: 0}))
	fmt.Println(g.Includes(Position{Row: -1, Column: 0}))

	// Output:
	// true
	// false
}

func ExamplePosition_FirstColumn() {
	g := Grid[int]{Values: [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}

	c1 := g.Cell(Position{Row: 0, Column: 0})
	c2 := g.Cell(Position{Row: 2, Column: 0})
	c3 := g.Cell(Position{Row: 1, Column: 1})

	fmt.Println(c1.Position.FirstColumn())
	fmt.Println(c2.Position.FirstColumn())
	fmt.Println(c3.Position.FirstColumn())

	// Output:
	// true
	// true
	// false
}

func ExampleGrid_Value() {
	g := Grid[int]{Values: [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}

	fmt.Println(g.Value(Position{Row: 0, Column: 0}))
	fmt.Println(g.Value(Position{Row: 1, Column: 1}))
	fmt.Println(g.Value(Position{Row: 2, Column: 2}))

	// Output:
	// 1
	// 5
	// 9
}

func ExampleGrid_Cell() {
	g := Grid[int]{Values: [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}

	fmt.Println(g.Cell(Position{Row: 0, Column: 0}))
	fmt.Println(g.Cell(Position{Row: 1, Column: 1}))
	fmt.Println(g.Cell(Position{Row: 2, Column: 2}))

	// Output:
	// {1 {0 0}}
	// {5 {1 1}}
	// {9 {2 2}}
}

func ExampleGrid_Size() {
	g := Grid[int]{Values: [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}

	// Number of columns is indicated by the number of columns in first row only
	nonuniform := Grid[int]{Values: [][]int{
		{1, 2, 3},
		{4, 5, 6, 7},
		{8, 9},
	}}

	empty := Grid[int]{}

	fmt.Println(g.Size())
	fmt.Println(nonuniform.Size())
	fmt.Println(empty.Size())

	// Output:
	// {3 3}
	// {3 3}
	// {0 0}
}

func TestGrid_AppendRow(t *testing.T) {
	t.Run("appends rows", func(t *testing.T) {
		g := Grid[int]{}
		g.AppendRow([]int{1, 2, 3})
		g.AppendRow([]int{5, 6, 7})

		want := "[[1 2 3] [5 6 7]]"
		got := fmt.Sprint(g.Rows())

		if got != want {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})

	t.Run("panic if rows size mismatch", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected a panic")
			}
		}()

		g := Grid[int]{}
		g.AppendRow([]int{1, 2, 3})
		g.AppendRow([]int{5, 6, 7, 8})
	})
}

package main

import (
	"fmt"
	"strings"

	"github.com/ahodieb/advent-of-code/common/slice"
	"github.com/ahodieb/brute/ansi"
	"github.com/ahodieb/brute/input"
)

func main() {
	in := input.FromPath("day12/input.txt")
	defer in.Close()

	var grid []string
	var start Pos
	var end Pos

	for i := 0; in.Scan(); i++ {
		if index := strings.Index(in.Text(), "E"); index != -1 {
			end = Pos{Column: index, Row: i}
		}

		if index := strings.Index(in.Text(), "S"); index != -1 {
			start = Pos{Column: index, Row: i}
		}

		t := in.Text()
		t = strings.Replace(t, "S", "a", 1)
		t = strings.Replace(t, "E", "z", 1)
		grid = append(grid, t)
	}

	//path := FindPath(grid, start, end)
	//fmt.Println(len(path), path)
	//
	//start = Pos{}
	//end = Pos{Column: 9}
	//grid = []string{"abcdefghijk"}
	//
	//start = Pos{Column: 9}
	//end = Pos{}
	//grid = []string{"lkjihgfedcba"}
	//
	//start = Pos{}
	//end = Pos{Row: 9, Column: 1}
	//grid = []string{
	//	"az",
	//	"bz",
	//	"cz",
	//	"dz",
	//	"ez",
	//	"fz",
	//	"gz",
	//	"hz",
	//	"iz",
	//	"jk",
	//}

	//path := FindPath(grid, start, end)
	//fmt.Println(len(path))
	//fmt.Println(
	//	VisualizeGrid(path, start, end, len(grid), len(grid[0])),
	//)

	startingPositions := FindAllStartPositions(grid)

	distances := BellmanFord(grid, start)
	fmt.Println(distances[end])

	fmt.Println(len(startingPositions))

	min := distances[end]
	for i, p := range startingPositions {
		fmt.Printf("Trying %v [%d/%d]\n", p, i+1, len(startingPositions))
		d := BellmanFord(grid, p)

		if dist, found := d[end]; found && dist < min {
			min = dist
		}
	}

	fmt.Println(min)
}

func FindAllStartPositions(grid []string) []Pos {
	var nodes []Pos
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 'a' {
				nodes = append(nodes, Pos{Row: r, Column: c})
			}
		}
	}

	return nodes
}

// BellmanFord not sure if this is actually BellmanFord's algorithm
// I think it is more Dijkstra's, I mixed things up trying to debug many silly bugs
func BellmanFord(graph []string, start Pos) map[Pos]int {
	distances := make(map[Pos]int)
	distances[start] = 0
	nodes := []Pos{start}

	for len(nodes) > 0 {
		cur := nodes[0]
		curDist := distances[cur]
		nodes = nodes[1:]

		adjacent := cur.GetAdjacentNodes()

		adjacent = slice.Filter(adjacent,
			WithinBounds(graph),
			CanClimb(graph, cur),
		)

		for _, n := range adjacent {
			_, ok := distances[n]
			if !ok || curDist+1 < distances[n] {
				distances[n] = curDist + 1
			}

			if !ok {
				nodes = append(nodes, n)
			}
		}

	}

	return distances
}

func CanClimb(grid []string, from Pos) func(Pos) bool {
	return func(to Pos) bool {
		return canClimb(grid[from.Row][from.Column], grid[to.Row][to.Column])
	}
}

func canClimb(from, to uint8) bool {
	return from == to || from+1 == to || to < from
}

func WithinBounds(grid []string) func(Pos) bool {
	return func(d Pos) bool {
		return d.Row >= 0 && d.Row < len(grid) && d.Column >= 0 && d.Column < len(grid[d.Row])
	}
}

type Pos struct {
	Row    int
	Column int
}

func (p *Pos) GetAdjacentNodes() []Pos {
	return []Pos{
		{Row: p.Row + 1, Column: p.Column},
		{Row: p.Row - 1, Column: p.Column},
		{Row: p.Row, Column: p.Column + 1},
		{Row: p.Row, Column: p.Column - 1},
	}
}

func VisualizeGrid(path []Pos, start, end Pos, h, w int) string {
	grid := make([][]string, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]string, w)
	}

	for i := 1; i < len(path); i++ {
		curr := path[i]
		prev := path[i-1]
		grid[prev.Row][prev.Column] = direction(prev, curr)
	}

	grid[start.Row][start.Column] = "S"
	grid[end.Row][end.Column] = "E"

	var sb strings.Builder
	for r := range grid {
		for c := range grid[r] {
			cell := grid[r][c]

			if cell == "" {
				cell = "."
			}
			_, _ = fmt.Fprintf(&sb, ansi.WhiteBG(fmt.Sprintf(" %s ", cell)))
		}
		_, _ = fmt.Fprintln(&sb)
	}

	return sb.String()
}

func direction(from Pos, to Pos) string {
	if from.Column == to.Column && from.Row < to.Row {
		return "v"
	}

	if from.Column == to.Column && from.Row > to.Row {
		return "^"
	}

	if from.Row == to.Row && from.Column < to.Column {
		return ">"
	}

	if from.Row == to.Row && from.Column > to.Column {
		return "<"
	}

	return fmt.Sprint(from, to)
}

package main

import (
	"fmt"
	"github.com/ahodieb/advent-of-code/common/grid"
	"github.com/ahodieb/advent-of-code/common/slice"
	"github.com/ahodieb/brute/ansi"
	"github.com/ahodieb/brute/input"
	"strings"
)

var digits = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func IsDigit(s string) bool {
	_, isDigit := digits[s]
	return isDigit
}

func IsSymbol(cell grid.Cell[string]) bool {
	return cell.Value != "." && !IsDigit(cell.Value)
}

type Part struct {
	Start  grid.Position
	End    grid.Position
	Number int
}

type Engine struct {
	G     grid.Grid[string]
	Parts []Part
}

func (e *Engine) Append(r string) {
	e.G.AppendRow(strings.Split(r, ""))
	e.Parts = append(e.Parts, Parse(r, e.G.Size().Rows-1)...)
}

func (e *Engine) ValidParts() []Part {
	return slice.Filter(e.Parts, func(p Part) bool { return p.IsValid(&e.G) })
}

func (e *Engine) SumValidParts() int {
	sum := 0
	for _, p := range e.ValidParts() {
		sum += p.Number
	}
	return sum
}

func (e *Engine) GearsRatios() int {
	sum := 0
	for i := e.G.Iterator(); i.Next(); {
		if i.Value() != "*" {
			continue
		}

		var adjacent []Part
		for _, part := range e.ValidParts() {
			for _, b := range part.Boundary() {
				if i.Position() == b {
					adjacent = append(adjacent, part)
					break
				}
			}
		}

		if len(adjacent) == 2 {
			sum += adjacent[0].Number * adjacent[1].Number
		}
	}

	return sum
}

func (p *Part) IsValid(g *grid.Grid[string]) bool {
	for _, b := range p.Boundary() {
		if !g.Includes(b) {
			continue
		}

		if IsSymbol(g.Cell(b)) {
			return true
		}
	}

	return false
}

func (p *Part) Boundary() []grid.Position {
	var boundary []grid.Position
	for i := p.Start.Column - 1; i <= p.End.Column+1; i++ {
		boundary = append(boundary,
			grid.Position{Row: p.Start.Row + 1, Column: i},
			grid.Position{Row: p.Start.Row, Column: i},
			grid.Position{Row: p.Start.Row - 1, Column: i},
		)
	}

	return boundary
}

func Parse(s string, row int) []Part {
	var parts []Part
	pos := grid.Position{Row: row, Column: 0}

	currentPart := Part{}
	inPart := false
	for _, r := range s {
		if IsDigit(string(r)) {
			if !inPart { // New part found
				inPart = true
				currentPart.Start = pos
				currentPart.End = pos
				currentPart.Number = digits[string(r)]
			} else { // still within part
				currentPart.End = pos
				currentPart.Number = currentPart.Number*10 + digits[string(r)]
			}
		} else {
			if inPart {
				parts = append(parts, currentPart)
				currentPart = Part{}
			}
			inPart = false
		}

		pos.Column++
	}

	if inPart {
		parts = append(parts, currentPart)
	}

	return parts
}

func (e *Engine) Render() string {
	r := grid.NewAsciiRenderer()

	// r.FmtPositionsAnsi(ansi.PurpleBG, e.Parts[0].Boundary()...)

	for _, p := range e.ValidParts() {
		r.FmtPositionsAnsi(ansi.GreenBG, p.Start)
		r.FmtPositionsAnsi(ansi.YellowBG, p.End)

	}

	r.FmtCell(IsSymbol, grid.AnsiFormatter(ansi.BlueBG))
	return r.Render(&e.G)
}

func main() {
	in := input.FromPath("2023/day3/input-small.txt")
	defer in.Close()

	e := Engine{}
	for in.Scan() {
		e.Append(in.Text())
	}

	fmt.Println(e.Render())
	fmt.Println(e.SumValidParts())
	fmt.Println(e.GearsRatios())
}

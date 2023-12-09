package grid

import (
	"fmt"
	"github.com/ahodieb/brute/ansi"
	"strings"
)

type Formatter func(cell Cell[string]) string

type ConditionalFormatter struct {
	Condition func(Cell[string]) bool
	Formatter Formatter
}

type AsciiRenderer struct {
	defaultFmt Formatter
	posFmt     map[Position]Formatter
	condFmt    []ConditionalFormatter
}

func NewAsciiRenderer() *AsciiRenderer {
	return &AsciiRenderer{
		defaultFmt: AnsiFormatter(ansi.WhiteBG),
		posFmt:     make(map[Position]Formatter),
	}
}

func (r *AsciiRenderer) SetDefaultFmt(f Formatter) *AsciiRenderer {
	r.defaultFmt = f
	return r
}

func AnsiFormatter(f func(string) string) Formatter {
	return func(cell Cell[string]) string {
		return f(fmt.Sprintf(" %s ", cell.Value))
	}
}

func (r *AsciiRenderer) FmtPositions(f Formatter, p ...Position) *AsciiRenderer {
	for _, pp := range p {
		r.posFmt[pp] = f
	}

	return r
}

func (r *AsciiRenderer) FmtPositionsAnsi(f func(string) string, p ...Position) *AsciiRenderer {
	return r.FmtPositions(AnsiFormatter(f), p...)
}

func (r *AsciiRenderer) FmtCell(condition func(Cell[string]) bool, formatter Formatter) *AsciiRenderer {
	r.condFmt = append(r.condFmt, ConditionalFormatter{Condition: condition, Formatter: formatter})
	return r
}

func (r *AsciiRenderer) Render(g *Grid[string]) string {
	var sb strings.Builder
	for i := g.Iterator(); i.Next(); {
		cell := i.Cell()

		if cell.Position.FirstColumn() {
			_, _ = fmt.Fprintln(&sb)
		}

		formatter := r.defaultFmt

		if f, found := r.posFmt[cell.Position]; found {
			formatter = f
		}

		for _, cf := range r.condFmt {
			if cf.Condition(cell) {
				formatter = cf.Formatter
			}
		}

		_, _ = fmt.Fprint(&sb, formatter(cell))
	}

	return sb.String()
}

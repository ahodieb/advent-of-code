package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/ahodieb/advent-of-code/common/slice"
	"github.com/ahodieb/brute/ansi"
)

type Point struct {
	x, y int
}

type Line struct {
	start Point
	end   Point
}

func (l *Line) IsHorizontal() bool {
	return l.start.x == l.end.x
}

func (l *Line) IsVertical() bool {
	return l.start.y == l.end.y
}

func (l *Line) IsDiagonal() bool {
	return math.Abs(float64(l.start.x-l.end.x)) == math.Abs(float64(l.start.y-l.end.y))
}

func (l *Line) DiagLen() int {
	s := slice.Min(l.start.x, l.end.x)
	e := slice.Max(l.start.x, l.end.x)
	return e - s
}

type Screen struct {
	pixels [][]int
}

func NewScreen(size int) Screen {
	p := make([][]int, size)
	for i := 0; i < size; i++ {
		p[i] = make([]int, size)
	}

	return Screen{pixels: p}
}

func (s *Screen) Draw(line Line) {
	if line.IsHorizontal() || line.IsVertical() {
		s.DrawHV(line)
	}

	if line.IsDiagonal() {
		s.DrawDiag(line)
	}
}

func (s *Screen) DrawHV(line Line) {
	sx := slice.Min(line.start.x, line.end.x)
	ex := slice.Max(line.start.x, line.end.x)
	sy := slice.Min(line.start.y, line.end.y)
	ey := slice.Max(line.start.y, line.end.y)

	for i := sx; i <= ex; i++ {
		for j := sy; j <= ey; j++ {
			s.pixels[j][i] += 1
		}
	}
}

func (s *Screen) DrawDiag(line Line) {
	var start, end Point

	if line.start.x < line.end.x {
		start = line.start
		end = line.end
	} else {
		start = line.end
		end = line.start
	}

	if start.y > end.y {
		for i := 0; i <= line.DiagLen(); i++ {
			s.pixels[start.x+i][start.y-i] += 1
		}
	} else {
		for i := 0; i <= line.DiagLen(); i++ {
			s.pixels[start.x+i][start.y+i] += 1
		}
	}

}

func (s *Screen) Count2s() int {
	count := 0
	size := len(s.pixels)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			pixel := s.pixels[i][j]
			if pixel > 1 {
				count += 1
			}

		}
	}
	return count
}

func (s *Screen) String() string {
	var sb strings.Builder
	size := len(s.pixels)
	for i := 0; i < size; i++ {
		if i == 0 {
			fmt.Fprint(&sb, ansi.RedBG("   "))
			for j := 0; j < size; j++ {
				fmt.Fprint(&sb, ansi.RedBG(fmt.Sprintf("%2d ", j)))
			}
			fmt.Fprintln(&sb)
		}
		for j := 0; j < size; j++ {
			if j == 0 {
				fmt.Fprint(&sb, ansi.RedBG(fmt.Sprintf("%2d ", i)))
			}
			pixel := s.pixels[i][j]
			fmtPixel := fmt.Sprintf("%2d ", pixel)
			if pixel > 1 {
				fmt.Fprint(&sb, ansi.Red(fmtPixel))
			} else if pixel > 0 {
				fmt.Fprint(&sb, ansi.Green(fmtPixel))
			} else {
				fmt.Fprint(&sb, fmtPixel)
			}
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}

func main() {
	f, err := os.Open("2021/day5/input-small.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var lines []Line
	for s.Scan() {
		var x1, y1, x2, y2 int
		fmt.Sscanf(s.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		lines = append(lines, Line{
			start: Point{x: x1, y: y1},
			end:   Point{x: x2, y: y2},
		})
	}
	screen := NewScreen(10)

	for _, l := range lines {
		screen.Draw(l)

		fmt.Println(l, l.IsDiagonal())
		fmt.Println(screen.String())
	}

	fmt.Println(screen.Count2s())
}

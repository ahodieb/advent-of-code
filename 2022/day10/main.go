package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahodieb/brute/ansi"
	"github.com/ahodieb/brute/input"
)

func main() {
	in := input.FromPath("day10/input.txt")
	defer in.Close()

	cpu := CPU{X: 1}
	lcd := Display{}
	raster := 20
	sum := 0

	for clock := 1; true; clock++ {
		cpu.Tick()
		lcd.Draw(cpu.X)

		if clock == raster { // Part1
			raster += 40
			strength := cpu.X * clock
			fmt.Println(ansi.Green(fmt.Sprintf("clock %03d x=%03d signal=%05d", clock, cpu.X, strength)))
			sum += strength
		}

		if cpu.Busy() {
			continue
		}

		if !in.Scan() {
			break
		}

		op := strings.Split(in.Text(), " ")
		if op[0] == "addx" {
			v, _ := strconv.Atoi(op[1])
			cpu.AddX(v)
		}
	}

	fmt.Println("Sum of signals", sum)
	fmt.Print(lcd.String())
}

type CPU struct {
	X      int
	buffer []int
}

func (c *CPU) Tick() {
	if len(c.buffer) > 0 {
		c.readFromBuffer()
	}
}

func (c *CPU) readFromBuffer() {
	c.X += c.buffer[0]
	c.buffer = c.buffer[1:]
}

func (c *CPU) AddX(x int) {
	c.buffer = append(c.buffer, 0, x)
}

func (c *CPU) Busy() bool {
	return len(c.buffer) != 0
}

type Display struct {
	sb strings.Builder
	p  int
}

func (d *Display) Draw(x int) {
	if d.p >= 40 {
		d.sb.WriteString("\n")
		d.p = 0
	}
	if d.p == x || d.p == x-1 || d.p == x+1 {
		d.sb.WriteString(ansi.RedBG(" "))
	} else {
		d.sb.WriteString(ansi.WhiteBG(" "))
	}
	d.p++
}

func (d *Display) String() string {
	return d.sb.String()
}

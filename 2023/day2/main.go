package main

import (
	"fmt"
	"github.com/ahodieb/advent-of-code/common/input"
	"strconv"
	"strings"
)

func extractDigits(v string) []int {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,

		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var digits []int
	for s := 0; s < len(v); s += 1 {
		for e := s + 1; e < len(v)+1; e++ {
			if d, isDigit := m[v[s:e]]; isDigit {
				digits = append(digits, d)
				break
			}
		}
	}
	return digits
}

type Cubes struct {
	Red   int
	Green int
	Blue  int
}

func (c Cubes) Power() int {
	return c.Red * c.Green * c.Blue
}

type Game struct {
	ID      int
	Samples []Cubes
}

func (g *Game) Possible(c Cubes) bool {
	for _, s := range g.Samples {
		if s.Red > c.Red || s.Blue > c.Blue || s.Green > c.Green {
			return false
		}
	}
	return true
}

func (g *Game) Min() Cubes {
	c := g.Samples[0]
	for _, s := range g.Samples[1:] {
		if s.Red > c.Red {
			c.Red = s.Red
		}

		if s.Green > c.Green {
			c.Green = s.Green
		}

		if s.Blue > c.Blue {
			c.Blue = s.Blue
		}
	}
	return c
}

func ParseGame(s string) Game {
	i := strings.Index(s, ":")
	id, _ := strconv.Atoi(s[5:i])
	samplesStr := s[i+1:]
	var samples []Cubes

	for _, sample := range strings.Split(strings.TrimSpace(samplesStr), ";") {
		var c Cubes
		for _, p := range strings.Split(sample, ",") {
			pair := strings.Split(strings.TrimSpace(p), " ")
			count, err := strconv.Atoi(pair[0])
			if err != nil {
				panic(fmt.Sprintf("%q failed %s", pair[0], err))
			}
			switch strings.TrimSpace(pair[1]) {
			case "red":
				c.Red = count
			case "blue":
				c.Blue = count
			case "green":
				c.Green = count
			}
		}

		samples = append(samples, c)
	}

	return Game{
		ID:      id,
		Samples: samples,
	}
}

func main() {
	in, err := input.FromPath("2023/day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	config := Cubes{Red: 12, Green: 13, Blue: 14}
	sum := 0
	sumPowers := 0
	for in.Scan() {
		game := ParseGame(in.Text())

		if game.Possible(config) {
			sum += game.ID
		}

		sumPowers += game.Min().Power()
	}

	fmt.Println(sum, sumPowers)
}

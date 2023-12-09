package main

import (
	"fmt"
	"github.com/ahodieb/advent-of-code/common/input"
	"github.com/ahodieb/advent-of-code/common/ints"
	"github.com/ahodieb/advent-of-code/common/slice"
	"math"
	"strings"
)

type Card struct {
	Name    string
	Wining  map[int]struct{}
	Numbers []int
}

func (c *Card) Points() int {
	count := 0
	for _, n := range c.Numbers {
		if _, found := c.Wining[n]; found {
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return int(math.Pow(2, float64(count-1)))
}

func ParseCard(s string) Card {
	card := strings.Split(s, ":")
	values := strings.Split(card[1], " | ")

	return Card{
		Name:    card[0],
		Wining:  slice.ToSet(ints.FromSpaceSeperated(values[0])...),
		Numbers: ints.FromSpaceSeperated(values[1]),
	}
}

func main() {
	in, err := input.FromArgsOrPath("2023/day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	sum := 0
	for in.Scan() {
		card := ParseCard(in.Text())
		sum += card.Points()
		//fmt.Println(card.Name, ":", card.Points(), " | ", card.Numbers)
	}

	fmt.Println(sum)
}

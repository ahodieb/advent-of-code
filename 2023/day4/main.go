package main

import (
	"fmt"
	"github.com/ahodieb/advent-of-code/common/input"
	"math"
)

type Card struct {
	Index   int
	Matches int
}

func (c *Card) Points() int {
	matches := c.Matches
	if matches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matches-1)))
}

func ParseCard(s string) Card {
	inNumber := false
	number := 0
	index := 0
	inIndex := true
	inWinning := false
	wining := make(map[int]struct{})
	matches := 0

	for _, r := range s[4:] {
		if '0' <= r && r <= '9' {
			if !inNumber {
				number = int(r - '0')
				inNumber = true
			} else {
				number = number*10 + int(r-'0')
			}
		} else {
			if inNumber {
				if inIndex {
					index = number
				} else if inWinning {
					wining[number] = struct{}{}
				} else {
					if _, found := wining[number]; found {
						matches += 1
					}
				}
			}

			inNumber = false
			number = 0
		}

		if r == ':' {
			inIndex = false
			inWinning = true
		}

		if r == '|' {
			inWinning = false
		}
	}

	if _, found := wining[number]; inNumber && found {
		matches += 1
	}

	return Card{
		Index:   index - 1,
		Matches: matches,
	}
}

func main() {
	in, err := input.FromArgsOrPath("2023/day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	var cards []Card
	var queue []Card
	var processed int
	var points = 0

	for in.Scan() {
		card := ParseCard(in.Text())
		cards = append(cards, card)
		points += card.Points()
	}

	queue = append(queue, cards...)
	for ; len(queue) > 0; queue = queue[1:] {
		card := queue[0]
		processed += 1
		queue = append(queue, cards[card.Index+1:][:card.Matches]...)
	}

	fmt.Println("Points:", points)
	fmt.Println("Processed:", processed)
}

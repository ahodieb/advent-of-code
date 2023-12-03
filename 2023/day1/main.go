package main

import (
	"fmt"
	"github.com/ahodieb/advent-of-code/common/input"
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

func main() {
	in, err := input.FromPath("2023/day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	sum := 0
	for in.Scan() {
		digits := extractDigits(in.Text())
		n := digits[0]*10 + digits[len(digits)-1]
		sum += n
	}

	fmt.Println(sum)
}

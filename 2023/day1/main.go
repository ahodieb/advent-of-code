package main

import (
	"fmt"
	"github.com/ahodieb/advent-of-code/common/input"
	"github.com/ahodieb/advent-of-code/common/slice"
	"strconv"
)

func filterDigits(s string) string {
	return string(slice.Filter([]rune(s), func(r rune) bool {
		digits := map[rune]struct{}{
			'0': {}, '1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {},
		}
		_, isDigit := digits[r]
		return isDigit
	}))
}

func main() {
	in, err := input.FromPath("2023/day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	sum := 0
	for in.Scan() {
		digits := filterDigits(in.Text())
		n, _ := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
		sum += n
	}

	fmt.Println(sum)
}

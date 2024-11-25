package main

import (
	"fmt"
	"sort"

	"github.com/ahodieb/brute/input"
	"github.com/ahodieb/brute/slices"
)

func main() {
	in := input.FromPath("2022/day1/input.txt")
	defer in.Close()

	sum := 0
	var sums []int

	for in.Scan() {
		if in.Text() != "" {
			sum += in.Int()
			continue
		}

		sums = append(sums, sum)
		sum = 0
	}
	sums = append(sums, sum)

	sort.Ints(sums)
	fmt.Println(sums[len(sums)-3:], slices.Sum(sums[len(sums)-3:]...))
}

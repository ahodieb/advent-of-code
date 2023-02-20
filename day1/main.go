package main

import (
	"fmt"
	"sort"

	"github.com/ahodieb/advent-of-code-2022/common/input"
	"github.com/ahodieb/advent-of-code-2022/common/slice"
)

func main() {
	in, err := input.FromPath("day1/input.txt")
	if err != nil {
		panic(err)
	}
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
	fmt.Println(sums[len(sums)-3:], slice.Sum(sums[len(sums)-3:]...))
}

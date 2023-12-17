package main

import (
	"fmt"
	"github.com/ahodieb/advent-of-code/common/input"
)

func ways(time, distance int) int {
	c := 0
	for wait := 0; wait < time; wait++ {
		if (time-wait)*wait > distance {
			c += 1
		}
	}

	return c
}
func main() {
	in, err := input.FromArgsOrPath("2023/day6/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	times := in.ReadPrefixedNumbers("Time:")
	distances := in.ReadPrefixedNumbers("Distance:")
	margin := 1
	for i := range times {
		margin = margin * ways(times[i], distances[i])
	}

	fmt.Println(margin)
}

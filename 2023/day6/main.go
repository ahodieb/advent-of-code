package main

import (
	"fmt"
	"github.com/ahodieb/advent-of-code/common/input"
	"strconv"
	"strings"
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

	t := strings.ReplaceAll(in.ReadPrefixedLine("Time:"), " ", "")
	d := strings.ReplaceAll(in.ReadPrefixedLine("Distance:"), " ", "")

	time, _ := strconv.Atoi(t)
	distance, _ := strconv.Atoi(d)

	fmt.Println(ways(time, distance))
}

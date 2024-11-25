package main

import (
	"fmt"
	"github.com/ahodieb/brute/input"
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
	in := input.FromPath("2023/day6/input.txt")
	defer in.Close()

	t := strings.ReplaceAll(strings.TrimPrefix(in.ReadText(), "Time:"), " ", "")
	d := strings.ReplaceAll(strings.TrimPrefix(in.ReadText(), "Distance:"), " ", "")

	time, _ := strconv.Atoi(t)
	distance, _ := strconv.Atoi(d)

	fmt.Println(ways(time, distance))
}

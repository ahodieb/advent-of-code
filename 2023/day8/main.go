package main

import (
	"fmt"
	"github.com/ahodieb/advent-of-code/common/input"
	"strings"
)

func main() {
	in, err := input.FromArgsOrPath("2023/day8/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	directions := in.ReadLine()
	in.ReadLine()

	m := make(map[string][]string)
	var current []string
	for in.Scan() {
		lr := strings.Split(in.Text()[7:len(in.Text())-1], ",")
		s := in.Text()[:3]
		m[s] = []string{lr[0], lr[1][1:]}
		if strings.HasSuffix(s, "A") {
			current = append(current, s)
		}
	}

	steps := make([]int, len(current))
	for i := range current {
		d := 0
		for !strings.HasSuffix(current[i], "Z") {
			n := 0
			if directions[d] == 'R' {
				n = 1
			}
			current[i] = m[current[i]][n]
			d = (d + 1) % len(directions)
			steps[i] += 1
		}
	}

	fmt.Println(steps) // LCM
}

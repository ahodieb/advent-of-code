package main

import (
	"fmt"
	"github.com/ahodieb/brute/input"
	"strings"
)

func main() {
	in := input.FromPath("2023/day8/input.txt")
	defer in.Close()

	directions := in.ReadText()
	in.ReadText()

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

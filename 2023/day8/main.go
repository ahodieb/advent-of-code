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
	for in.Scan() {
		lr := strings.Split(in.Text()[7:len(in.Text())-1], ",")
		m[in.Text()[:3]] = []string{lr[0], lr[1][1:]}
	}

	c := "AAA"
	d := 0
	s := 0
	for c != "ZZZ" {
		n := 0
		if directions[d] == 'R' {
			n = 1
		}

		c = m[c][n]
		d = (d + 1) % len(directions)
		s += 1
	}

	fmt.Println(s)
}

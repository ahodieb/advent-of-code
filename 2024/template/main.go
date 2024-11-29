package main

import (
	"fmt"
	"github.com/ahodieb/brute/input"
)

func main() {
	day := "template"
	p1(fmt.Sprintf("2024/%s/part1-small.txt", day))
	p1(fmt.Sprintf("2024/%s/part1.txt", day))
	p2(fmt.Sprintf("2024/%s/part2-small.txt", day))
	p2(fmt.Sprintf("2024/%s/part2.txt", day))
}

func p1(p string) {
	in := input.FromPath(p)
	defer in.Close()

	sum := 0
	for in.Scan() {
		sum += in.Int()
	}

	fmt.Println(sum)
}

func p2(p string) {
	p1(p)
}

package main

import (
	"fmt"
	"github.com/ahodieb/brute/input"
)

func main() {
	day := "day2"
	solve(fmt.Sprintf("2024/%s/input-small.txt", day))
	solve(fmt.Sprintf("2024/%s/input.txt", day))
}

func solve(p string) {
	in := input.FromPath(p)
	defer in.Close()
	for in.Scan() {
	}
}

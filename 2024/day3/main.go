package main

import (
	"fmt"
	"github.com/ahodieb/brute/input"
	"regexp"
	"strconv"
)

func main() {
	day := "day3"
	solve(fmt.Sprintf("2024/%s/input-small.txt", day)) // 161
	solve(fmt.Sprintf("2024/%s/input.txt", day))       // 174561379
}

func solve(p string) {
	in := input.FromPath(p)
	defer in.Close()
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	sum := 0
	for in.Scan() {
		parts := re.FindAllString(in.Text(), -1)
		for _, p := range parts {
			groups := re.FindStringSubmatch(p)
			a, _ := strconv.Atoi(groups[1])
			b, _ := strconv.Atoi(groups[2])
			sum = sum + (a * b)
			//fmt.Println(a, b, a*b)
		}
	}

	fmt.Println(sum)
}

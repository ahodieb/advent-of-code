package main

import (
	"fmt"
	"github.com/ahodieb/brute/input"
	"slices"
)

func main() {
	day := "day2"
	solve(fmt.Sprintf("2024/%s/input-small.txt", day)) // 2 4
	solve(fmt.Sprintf("2024/%s/input.txt", day))       // 341 404
}

func solve(p string) {
	in := input.FromPath(p)
	defer in.Close()

	safe := 0
	safe2 := 0

	for in.Scan() {
		reports := in.Ints()
		if isSafe(reports) {
			safe++
		} else {
			for i := 0; i < len(reports); i++ {
				var dst []int
				dst = append(dst, reports[:i]...)
				dst = append(dst, reports[i+1:]...)
				if isSafe(dst) {
					safe2++
					break
				}
			}
		}
	}

	fmt.Println(safe, safe+safe2)
}

func isSafe(reports []int) bool {
	if len(reports) < 2 {
		return true
	}

	if reports[0] == reports[1] {
		return false
	}

	if reports[0] > reports[1] {
		slices.Reverse(reports)
	}

	for i := 1; i < len(reports); i++ {
		if reports[i-1] >= reports[i] {
			return false
		}

		if reports[i]-reports[i-1] > 3 {
			return false
		}
	}

	return true
}

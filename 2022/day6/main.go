package main

import (
	"fmt"

	"github.com/ahodieb/advent-of-code/common/input"
)

func main() {
	in, err := input.FromPath("2022/day6/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	for in.Scan() {
		fmt.Println(
			findMarker(in.Text(), 4),
			findMarker(in.Text(), 14),
		)
	}
}

func findMarker(signal string, size int) int {
	for i := size; i < len(signal); i++ {
		if allDiff(signal[i-size : i]) {
			return i
		}
	}

	return 0
}

func allDiff(marker string) bool {
	hash := make(map[uint8]struct{})

	for i := 0; i < len(marker); i++ {
		if _, ok := hash[marker[i]]; ok {
			return false
		}

		hash[marker[i]] = struct{}{}
	}

	return true
}

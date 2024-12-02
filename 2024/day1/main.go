package main

import (
	"fmt"
	"github.com/ahodieb/brute/histogram"
	"github.com/ahodieb/brute/input"
	"slices"
)

func main() {
	day := "day1"
	solve(fmt.Sprintf("2024/%s/input-small.txt", day)) // 11, 31
	solve(fmt.Sprintf("2024/%s/input.txt", day))       // 1938424, 22014209
}

func solve(p string) {
	in := input.FromPath(p)
	defer in.Close()

	var l1, l2 []int
	for in.Scan() {
		i := in.Ints()
		l1 = append(l1, i[0])
		l2 = append(l2, i[1])
	}

	h := histogram.FromSlice(l2)
	slices.Sort(l1)
	slices.Sort(l2)

	if len(l1) != len(l2) {
		panic(fmt.Errorf("lists don't have the same length %v, %v", len(l1), len(l2)))
	}

	dist := 0
	sim := 0
	for i := 0; i < len(l1); i++ {
		dist = dist + max(l1[i], l2[i]) - min(l1[i], l2[i])
		sim = sim + (l1[i] * h.Count(l1[i]))
	}

	fmt.Println("Distance:", dist)
	fmt.Println("Similarity:", sim)
}

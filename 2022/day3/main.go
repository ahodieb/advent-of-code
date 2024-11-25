package main

import (
	"fmt"
	ch "github.com/ahodieb/advent-of-code/common/input"
	"github.com/ahodieb/brute/input"
)

func main() {
	part1()
	part2()
}

func part2() {
	in := ch.ChunkedFromPath("day3/input.txt", 3)
	defer in.Close()

	sum := 0
	for in.Scan() {
		item := findCommon(in.Chunk()...)
		priority := findPriority(item)
		sum += priority
	}

	fmt.Println(sum)
}

func part1() {
	in := input.FromPath("2022/day3/input.txt")
	defer in.Close()

	sum := 0
	for in.Scan() {
		item := findItem(in.Text())
		priority := findPriority(item)
		sum += priority
	}

	fmt.Println(sum)
}

func findPriority(item uint8) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}

	if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}

	return 0
}

func findItem(items string) uint8 {
	return findCommon(items[:len(items)/2], items[len(items)/2:])
}

func findCommon(items ...string) uint8 {
	hash := make(map[uint8]struct{})

	for i := 0; i < len(items[0]); i++ {
		hash[items[0][i]] = struct{}{}
	}

	for _, bag := range items {
		intersection := make(map[uint8]struct{})

		for i := 0; i < len(bag); i++ {
			if _, ok := hash[bag[i]]; ok {
				intersection[bag[i]] = struct{}{}
			}
		}
		hash = intersection
	}

	for k := range hash {
		return k
	}

	return 0
}

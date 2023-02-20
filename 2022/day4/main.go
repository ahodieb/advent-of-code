package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahodieb/advent-of-code/common/input"
)

func main() {
	in, err := input.FromPath("2022/day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	countTotalOverlap := 0
	countPartialOverlap := 0
	for in.Scan() {
		if totalOverlap(in.Text()) {
			countTotalOverlap += 1
		}

		if partialOverlap(in.Text()) {
			countPartialOverlap += 1
		}
	}

	fmt.Println(countTotalOverlap, countPartialOverlap)
}

func totalOverlap(pair string) bool {
	ranges := strings.Split(pair, ",")
	p1 := FromString(ranges[0])
	p2 := FromString(ranges[1])

	return p1.Includes(p2) || p2.Includes(p1)
}

func partialOverlap(pair string) bool {
	ranges := strings.Split(pair, ",")
	p1 := FromString(ranges[0])
	p2 := FromString(ranges[1])

	return p1.Overlaps(p2) || p2.Overlaps(p1)
}

type Range struct {
	Start int
	End   int
}

func (r *Range) Includes(rr Range) bool {
	return r.Start <= rr.Start && r.End >= rr.End
}

func (r *Range) Overlaps(rr Range) bool {
	return r.Start <= rr.Start && r.End >= rr.Start || r.Start <= rr.End && r.End >= rr.End
}

func FromString(r string) Range {
	numbers := strings.Split(r, "-")
	start, _ := strconv.Atoi(numbers[0])
	end, _ := strconv.Atoi(numbers[1])
	return Range{Start: start, End: end}
}

package main

import (
	"fmt"
	"github.com/ahodieb/advent-of-code/common/ints"
	"github.com/ahodieb/advent-of-code/common/slice"
	"github.com/ahodieb/brute/input"
	"strings"
)

//
//type Parser[T any] struct {
//	text  string
//	pos   int
//	start int
//
//	initialState stateFn[T]
//	output       T
//}
//
//type stateFn[T any] func(p *Parser[T]) stateFn[T]
//
//func (p *Parser[T]) Parse() T {
//	for state := p.initialState; state != nil; {
//		state = state(p)
//	}
//
//	return p.output
//}
//
//func (p *Parser[T]) next(i int) {
//	p.pos += i
//}
//
//func (p *Parser[T]) hasNext() bool {
//	return p.pos < len(p.text)
//}
//
//func (p *Parser[T]) isDigit() bool {
//	return '0' <= p.text[p.pos] && p.text[p.pos] <= '9'
//}
//
//type ParserOutput struct {
//	Seeds []int
//}
//
//func initial(p *Parser[ParserOutput]) stateFn[ParserOutput] {
//	const seeds = "seeds: "
//	for ; p.hasNext(); p.next(1) {
//		if strings.HasSuffix(p.text[p.pos:], seeds) {
//			p.start = p.pos
//			return seedsState
//		}
//	}
//
//	return nil
//}
//
//func seedsState(p *Parser[ParserOutput]) stateFn[ParserOutput] {
//	for ; p.hasNext(); p.next(1) {
//		if p.isDigit() {
//
//		}
//	}
//
//	return nil
//}

type Mapping struct {
	Src      string
	Dest     string
	SrcStart int
	DesStart int
	RangeLen int
}

func (m Mapping) String() string {
	return fmt.Sprintf("{%s => %s [%d:%d * %d]}", m.Src, m.Dest, m.SrcStart, m.DesStart, m.RangeLen)
}

func ParseMapping(in *input.Input) []Mapping {
	var mappings []Mapping
	srcToDst := strings.Split(strings.TrimSuffix(in.Text(), "map:"), "-to-")
	for in.Scan() {
		if in.Text() == "" {
			break
		}

		ranges := ints.FromSpaceSeperated(in.Text())
		if len(ranges) != 3 {
			panic(fmt.Sprintf("expected 3 numbers, got: %v", ranges))
		}
		mappings = append(mappings, Mapping{
			Src:      srcToDst[0],
			Dest:     srcToDst[1],
			SrcStart: ranges[1],
			DesStart: ranges[0],
			RangeLen: ranges[2],
		})
	}

	return mappings
}

func main() {
	in := input.FromPath("2023/day5/input.txt")
	defer in.Close()

	in.Scan()
	seeds := ints.FromSpaceSeperated(strings.TrimPrefix(in.Text(), "seeds: "))
	seeds = expandSeeds(seeds)
	mapped := append([]int{}, seeds...)

	for in.Scan() {
		if strings.HasSuffix(in.Text(), "map:") {
			mappings := ParseMapping(in)
			for i := range mapped {
				v := mapped[i]
				for _, mm := range mappings {
					if mm.SrcStart <= v && v < mm.SrcStart+mm.RangeLen {
						offset := v - mm.SrcStart
						newV := mm.DesStart + offset
						mapped[i] = newV
					}
				}
			}
		}
	}

	//fmt.Println("Locations", seeds)
	fmt.Println("Closest:", slice.Min(mapped...))
}

func expandSeeds(seeds []int) []int {
	var expanded []int
	for i := 0; i < len(seeds); i += 2 {
		for s := seeds[i]; s < seeds[i]+seeds[i+1]; s++ {
			expanded = append(expanded, s)
		}
	}
	return expanded
}

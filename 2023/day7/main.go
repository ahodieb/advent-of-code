package main

import (
	"cmp"
	"fmt"
	"github.com/ahodieb/advent-of-code/common/input"
	"slices"
	"strconv"
	"strings"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var CardStrength = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type RuneCountPair struct {
	Rune  rune
	Count int
}

func (p RuneCountPair) String() string {
	return fmt.Sprintf("%c:%d", p.Rune, p.Count)
}

type Hand string

func (h Hand) Count() []RuneCountPair {
	counts := make(map[rune]int)
	var pairs []RuneCountPair
	for _, r := range h {
		counts[r] += 1
	}

	for r, c := range counts {
		pairs = append(pairs, RuneCountPair{Rune: r, Count: c})
	}

	slices.SortStableFunc(pairs, func(a, b RuneCountPair) int { return cmp.Compare(b.Count, a.Count) })
	return pairs
}

func (h Hand) Type() int {
	c := h.Count()

	switch {
	case c[0].Count == 5:
		return FiveOfAKind
	case c[0].Count == 4:
		return FourOfAKind
	case c[0].Count == 3 && c[1].Count == 2:
		return FullHouse
	case c[0].Count == 3:
		return ThreeOfAKind
	case c[0].Count == 2 && c[1].Count == 2:
		return TwoPair
	case c[0].Count == 2:
		return OnePair
	default:
		return HighCard
	}
}

func (h Hand) Cmp(hh Hand) int {
	if h.Type() != hh.Type() {
		return cmp.Compare(h.Type(), hh.Type())
	}

	for i := range h {
		if h[i] != hh[i] {
			return cmp.Compare(CardStrength[rune(h[i])], CardStrength[rune(hh[i])])
		}
	}

	return 0
}

type HandAndBid struct {
	Hand Hand
	Bid  int
}

func (h HandAndBid) String() string {
	return fmt.Sprintf("%s | %d | %v", h.Hand, h.Hand.Type(), h.Hand.Count())
}

func ParseLine(s string) HandAndBid {
	bid, err := strconv.Atoi(strings.TrimSpace(s[5:]))
	if err != nil {
		panic(fmt.Sprintf("failed to parse %q, %v", s, err))
	}
	return HandAndBid{Hand: Hand(s[:5]), Bid: bid}
}

func main() {
	in, err := input.FromArgsOrPath("2023/day7/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	hands := input.ParseLines(in, ParseLine)
	slices.SortStableFunc(hands, func(a, b HandAndBid) int {
		return a.Hand.Cmp(b.Hand)
	})

	total := 0
	for i := range hands {
		// fmt.Println(h, "|", h.Bid, "x", i+1, "=", (i+1)*hands[i].Bid)
		total += (i + 1) * hands[i].Bid
	}

	fmt.Println(total)
}

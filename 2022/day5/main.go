package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahodieb/brute/input"
)

func main() {
	in := input.FromPath("2022/day5/input.txt")
	defer in.Close()

	var rawStacks []string
	for in.Scan() {
		if in.Text() == "" {
			break
		}

		rawStacks = append(rawStacks, in.Text())
	}

	stacks := ToStacks(rawStacks[:len(rawStacks)-1])
	for in.Scan() {
		//stacks.Do(ToMove(in.Text()))
		stacks.Do9001(ToMove(in.Text()))
	}
	fmt.Println(stacks.Top())
}

type Stacks struct {
	Stacks [][]string
}

func (s *Stacks) Do(move Move) {
	for i := 0; i < move.Repeat; i++ {
		v := s.Stacks[move.From-1][0]
		s.Stacks[move.From-1] = s.Stacks[move.From-1][1:]
		s.Stacks[move.To-1] = append([]string{v}, s.Stacks[move.To-1]...)
	}
}

func (s *Stacks) Do9001(move Move) {
	var block []string
	for i := 0; i < move.Repeat; i++ {
		v := s.Stacks[move.From-1][0]
		block = append(block, v)
		s.Stacks[move.From-1] = s.Stacks[move.From-1][1:]
	}

	s.Stacks[move.To-1] = append(block, s.Stacks[move.To-1]...)
}

func (s *Stacks) Top() string {
	var sb strings.Builder
	for i := range s.Stacks {
		if len(s.Stacks[i]) > 0 {
			sb.WriteString(s.Stacks[i][0])
		}
	}

	return sb.String()
}

func ToStacks(raw []string) *Stacks {
	count := len(strings.Split(raw[len(raw)-1], " "))
	stacks := make([][]string, count)

	for i := len(raw) - 1; i >= 0; i-- {
		indx := 0
		nxt := 1
		for nxt < len(raw[i]) {
			c := string(raw[i][nxt])

			if strings.TrimSpace(c) != "" {
				stacks[indx] = append([]string{c}, stacks[indx]...)
			}

			nxt += 4
			indx += 1
		}

	}

	return &Stacks{Stacks: stacks}
}

type Move struct {
	From   int
	To     int
	Repeat int
}

func ToMove(move string) Move {
	repeat, _ := strconv.Atoi(strings.TrimSpace(
		move[(strings.Index(move, "move") + 4):strings.Index(move, "from")],
	))

	from, _ := strconv.Atoi(strings.TrimSpace(
		move[(strings.Index(move, "from") + 4):strings.Index(move, "to")],
	))

	to, _ := strconv.Atoi(strings.TrimSpace(
		move[(strings.Index(move, "to") + 2):],
	))

	return Move{From: from, To: to, Repeat: repeat}
}

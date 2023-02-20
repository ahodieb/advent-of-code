package main

import (
	"fmt"

	"github.com/ahodieb/advent-of-code-2022/common/input"
)

func main() {
	in, err := input.FromPath("2022/day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	var score Score
	for in.Scan() {
		score.Add(in.Text())
	}

	fmt.Println(score.String())
}

type Score struct {
	score  int
	score2 int
}

func (s *Score) Add(round string) {
	s.score += scoreForRound(round)
	s.score2 += scoreForRoundV2(round)
}

func scoreForRound(match string) int {
	score := 0

	switch match[2:] {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}

	if match == "A X" || match == "B Y" || match == "C Z" {
		score += 3
	}

	if match == "A Y" || match == "B Z" || match == "C X" {
		score += 6
	}

	return score
}

func scoreForRoundV2(round string) int {
	score := 0

	switch round[2:] {
	case "Y":
		score += 3
	case "Z":
		score += 6
	}

	if round == "A X" || round == "B Z" || round == "C Y" {
		score += 3
	}

	if round == "A Y" || round == "B X" || round == "C Z" {
		score += 1
	}

	if round == "A Z" || round == "B Y" || round == "C X" {
		score += 2
	}

	return score
}

func (s *Score) String() string {
	return fmt.Sprintf("Score: %d, New Score: %d", s.score, s.score2)
}

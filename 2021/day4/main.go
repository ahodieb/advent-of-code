package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	size         int
	cells        []int
	cellValueMap map[int]int
	sum          int
	marked       map[int]struct{}
	lastMarked   int
	won          bool
}

func NewBoard(size int) Board {
	return Board{
		size:         size,
		cellValueMap: make(map[int]int),
		marked:       make(map[int]struct{}),
	}
}
func (b *Board) Size() int {
	return b.size
}

func (b *Board) AddStrCell(s string) {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}

	b.AddCells(n)
}

func (b *Board) AddCells(row ...int) {
	start := len(b.cells)
	b.cells = append(b.cells, row...)

	for i, n := range row {
		b.sum += n
		b.cellValueMap[n] = i + start
	}
}

func (b *Board) Mark(numbers ...int) {
	if b.won {
		return
	}

	for _, n := range numbers {
		_, found := b.cellValueMap[n]
		if found {
			b.sum -= n
			b.marked[n] = struct{}{}
			b.lastMarked = n
		}
	}

	b.IsWon()
}

func (b *Board) IsWon() bool {
	if !b.won {
		b.won = b.hasCompleteRow() || b.hasCompleteColumn()
	}

	return b.won
}

func (b *Board) Score() int {
	return b.sum * b.lastMarked
}

func (b *Board) String() string {
	var sb strings.Builder
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			cell := b.cells[i*b.size+j]

			if b.isMarked(cell) {
				fmt.Fprint(&sb, boldCell(cell))
			} else {
				fmt.Fprintf(&sb, "%2d ", cell)
			}

		}
		fmt.Fprintln(&sb)
	}

	fmt.Fprintln(&sb)
	if b.IsWon() {
		fmt.Fprintf(&sb, "score: %d, ", b.Score())
	}
	fmt.Fprintf(&sb, "sum: %d, last-marked: %d\n", b.sum, b.lastMarked)
	return sb.String()
}

func boldCell(n int) string {
	return fmt.Sprintf("\u001b[1m\u001b[31m%2d \u001b[0m", n)
}

func (b *Board) hasCompleteRow() bool {
	for i := 0; i < b.size; i++ {
		complete := true

		for j := 0; j < b.size && complete; j++ {
			c := b.cells[i*5+j]
			_, marked := b.marked[c]
			complete = marked && complete
		}

		if complete {
			return true
		}
	}

	return false
}

func (b *Board) hasCompleteColumn() bool {
	for i := 0; i < b.size; i++ {
		complete := true

		for j := 0; j < b.size && complete; j++ {
			c := b.cells[j*5+i]
			complete = b.isMarked(c) && complete
		}

		if complete {
			return true
		}
	}

	return false
}

func (b *Board) isMarked(c int) bool {
	_, marked := b.marked[c]
	return marked
}

func main() {
	f, err := os.Open("2021/day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	numbers := scanNumbers(s)
	boards := scanBoards(s)

	var winning []*Board

	for _, n := range numbers {
		for _, b := range boards {
			if b.IsWon() {
				continue
			}

			b.Mark(n)

			if b.IsWon() {
				winning = append(winning, b)
			}
		}
	}

	// for _, b := range boards {
	// 	fmt.Println(b.String())
	// }

	fmt.Printf("Found %d winning boards out of %d\n", len(winning), len(boards))
	if len(winning) > 0 {
		fmt.Println(winning[0].String())
		fmt.Println(winning[len(winning)-1].String())
	} else {

		fmt.Println("❌ No winning board found")
	}
}

func scanBoards(s *bufio.Scanner) []*Board {
	const size = 5
	const rlen = size*3 - 1

	var boards []*Board
	for s.Scan() {
		b := NewBoard(size)

		for i := 0; i < size; i++ {
			r := s.Text()
			for j := 0; j < rlen; j += 3 {
				b.AddStrCell(r[j : j+2])
			}
			s.Scan()
		}
		boards = append(boards, &b)
	}

	return boards
}

func scanNumbers(s *bufio.Scanner) []int {
	s.Scan()
	defer s.Scan()

	cells := strings.Split(s.Text(), ",")
	var n []int
	for _, c := range cells {
		v, err := strconv.Atoi(strings.TrimSpace(c))
		if err != nil {
			panic(err)
		}

		n = append(n, v)
	}

	return n
}

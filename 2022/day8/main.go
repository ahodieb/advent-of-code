package main

import (
	"fmt"

	"github.com/ahodieb/advent-of-code/common/ansi"
	"github.com/ahodieb/advent-of-code/common/input"
)

func main() {
	in, err := input.FromPath("2022/day8/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	var grid [][]uint8
	for in.Scan() {
		var row []uint8
		for i := range in.Text() {
			row = append(row, in.Text()[i]-'0')
		}
		grid = append(grid, row)
	}

	maxScore := MaxScenicScore(grid)
	countVisibility := CountVisible(grid)

	for r := range grid {
		for c := range grid[r] {
			s := score(grid, r, c)
			cell := fmt.Sprintf(" %d ", grid[r][c])

			if s >= maxScore {
				cell = ansi.RedBG(cell)
			} else {
				cell = ansi.WhiteBG(cell)
			}

			fmt.Print(cell)
		}
		println()
	}

	fmt.Println("visible: ", countVisibility)
	fmt.Println("highest score: ", maxScore)

}

func MaxScenicScore(grid [][]uint8) int {
	max := 0
	for r := range grid {
		for c := range grid[r] {
			if s := score(grid, r, c); s > max {
				max = s
			}
		}
	}
	return max
}

func CountVisible(grid [][]uint8) int {
	count := 0
	for r := range grid {
		for c := range grid[r] {
			if visible(grid, r, c) {
				count += 1
			}
		}
	}
	return count
}

func visible(grid [][]uint8, r, c int) bool {
	vis := true
	for rr := r + 1; rr < len(grid); rr++ {
		if grid[rr][c] >= grid[r][c] {
			vis = false
			break
		}
	}

	if vis {
		return vis
	}

	vis = true
	for rr := r - 1; rr >= 0; rr-- {
		if grid[rr][c] >= grid[r][c] {
			vis = false
			break
		}
	}

	if vis {
		return vis
	}

	vis = true
	for cc := c + 1; cc < len(grid[r]); cc++ {
		if grid[r][cc] >= grid[r][c] {
			vis = false
			break
		}
	}

	if vis {
		return vis
	}

	vis = true
	for cc := c - 1; cc >= 0; cc-- {
		if grid[r][cc] >= grid[r][c] {
			vis = false
			break
		}
	}

	return vis
}

func score(grid [][]uint8, r, c int) int {
	right := 0
	for rr := r + 1; rr < len(grid); rr++ {
		right += 1
		if grid[rr][c] >= grid[r][c] {
			break
		}
	}

	left := 0
	for rr := r - 1; rr >= 0; rr-- {
		left += 1
		if grid[rr][c] >= grid[r][c] {
			break
		}
	}

	top := 0
	for cc := c + 1; cc < len(grid[r]); cc++ {
		top += 1
		if grid[r][cc] >= grid[r][c] {
			break
		}
	}

	bottom := 0
	for cc := c - 1; cc >= 0; cc-- {
		bottom += 1
		if grid[r][cc] >= grid[r][c] {
			break
		}
	}

	return right * left * bottom * top
}

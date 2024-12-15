package main

import (
	"fmt"
	"github.com/ahodieb/brute/input"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don't\(\))`)

func main() {
	day := "day3"
	solve(fmt.Sprintf("2024/%s/input-small.txt", day))    // 161
	solve(fmt.Sprintf("2024/%s/input-small-p2.txt", day)) // 48
	solve(fmt.Sprintf("2024/%s/input.txt", day))          // 174561379 106921067
}

func solve(p string) {
	in := input.FromPath(p)
	defer in.Close()

	sum := 0
  sum2 := 0
  enabled := true
	for in.Scan() {
		for _, m := range re.FindAllStringSubmatch(in.Text(), -1) {

			if strings.HasPrefix(m[0], "mul") {
				a, _ := strconv.Atoi(m[2])
				b, _ := strconv.Atoi(m[3])
        sum2 = sum2 + (a * b)
				if enabled {
					sum = sum + (a * b)
				}
			}

      if m[0] == "do()" {
        enabled = true
      }

      if m[0] == "don't()" {
        enabled = false
      }
		}
	}

	fmt.Println(sum2, sum)
}


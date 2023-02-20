package main

import (
	"fmt"
	"strconv"

	"github.com/ahodieb/advent-of-code/common/input"
)

func main() {
	in := input.MustFromPath("day13/input.txt")
	defer in.Close()

	var packets []string
	for in.Scan() {
		if in.Text() == "" {
			continue
		}

		packets = append(packets, in.Text())
	}

	count := 0
	for i := 0; i < len(packets); i += 2 {
		if InOrder(Parse(packets[i]), Parse(packets[i+1])) {
			count += i/2 + 1
		}
	}

	fmt.Println(count)
}

//func InOrder(left, right string) {
//	var iLeft []string
//	var iRight []string
//
//	curItem := ""
//
//}

func Parse(packet string) []string {
	var tree []string

	curItem := ""
	for _, c := range packet {
		if c == '[' {
			tree = append(tree, fmt.Sprintf("%c", c))
			curItem = ""
		} else if c == ']' {
			if curItem != "" {
				tree = append(tree, curItem)
			}
			tree = append(tree, fmt.Sprintf("%c", c))
			curItem = ""
		} else if c == ',' {
			if curItem != "" {
				tree = append(tree, curItem)
			}
			curItem = ""
		} else {
			curItem = fmt.Sprintf("%s%c", curItem, c)
		}
	}

	return tree
}

func InOrder(left, right []string) bool {
	for len(left) > 0 {
		if len(right) == 0 {
			fmt.Println("- Right side ran out of items, so inputs are not in the right order")
			return false
		}

		l := left[0]
		r := right[0]
		left = left[1:]
		right = right[1:]

		if l == r {
			continue
		}

		if IsInt(l) && IsInt(r) {
			return ToInt(l) < ToInt(r)
		}

		if IsInt(l) && r == "[" {
			left = append([]string{l, "]"}, left...)
			continue
		}
		if l == "[" && IsInt(r) {
			right = append([]string{r, "]"}, right...)
			continue
		}
	}
	return true
}

func IsInt(i string) bool {
	_, err := strconv.Atoi(i)
	return err == nil
}

func ToInt(i string) int {
	v, _ := strconv.Atoi(i)
	return v
}

//
//func InOrderInts(left, right string) bool {
//	leftInt, err := strconv.Atoi(left)
//	if err != nil {
//		panic(fmt.Errorf("left: %v is not an int", left))
//	}
//
//	rightInt, err := strconv.Atoi(left)
//	if err != nil {
//		panic(fmt.Errorf("right: %v is not an int", right))
//	}
//
//	return leftInt <= rightInt
//}

package ints

import (
	"fmt"
	"strconv"
	"strings"
)

func FromSpaceSeperated(s string) []int {
	var numbers []int
	for _, ss := range strings.Split(s, " ") {
		trimmed := strings.TrimSpace(ss)
		if trimmed == "" {
			continue
		}

		n, err := strconv.Atoi(ss)
		if err != nil {
			panic(fmt.Errorf("%q is not a number, %w", ss, err))
		}

		numbers = append(numbers, n)
	}

	return numbers
}

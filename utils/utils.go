package utils

import "fmt"

func Min(first int, rest ...int) int {
	if len(rest) == 0 {
		return first
	}

	m := first
	for _, n := range rest {
		if n < m {
			m = n
		}
	}
	return m
}

func Max(first int, rest ...int) int {
	if len(rest) == 0 {
		return first
	}

	m := first
	for _, n := range rest {
		if n > m {
			m = n
		}
	}
	return m
}

func AnsiRed(s string) string {
	return fmt.Sprintf("\u001b[31m%s\u001b[0m", s)
}

func AnsiGreen(s string) string {
	return fmt.Sprintf("\u001b[32m%s\u001b[0m", s)
}

func AnsiUnderline(s string) string {
	return fmt.Sprintf("\u001b[4m%s\u001b[0m", s)
}
func AnsiBGRed(s string) string {
	return fmt.Sprintf("\u001b[41m%s\u001b[0m", s)
}

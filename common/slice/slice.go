package slice

func Sum(numbers ...int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}

	return s
}

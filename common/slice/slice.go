package slice

func Sum(numbers ...int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}

	return s
}

func Min[T uint8 | int | int64 | float32 | float64](numbers ...T) T {
	var min T
	if len(numbers) == 0 {
		return min
	}

	min = numbers[0]
	for i := range numbers {
		if numbers[i] < min {
			min = numbers[i]
		}
	}

	return min
}

func Filter[T any](items []T, filters ...func(i T) bool) []T {
	var newSlice []T

	for i := range items {
		item := items[i]
		included := true
		for _, f := range filters {
			if !f(item) {
				included = false
				break
			}
		}

		if included {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}

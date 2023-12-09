package slice

type numbers interface {
	uint8 | int | int64 | float32 | float64
}

func Min[T numbers](numbers ...T) T {
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

func Max[T uint8 | int | int64 | float32 | float64](numbers ...T) T {
	var max T
	if len(numbers) == 0 {
		return max
	}

	max = numbers[0]
	for i := range numbers {
		if numbers[i] > max {
			max = numbers[i]
		}
	}

	return max
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

func ToSet[T comparable](items ...T) map[T]struct{} {
	set := make(map[T]struct{})
	for i := range items {
		set[items[i]] = struct{}{}
	}

	return set
}

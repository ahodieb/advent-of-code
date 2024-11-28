package slice

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

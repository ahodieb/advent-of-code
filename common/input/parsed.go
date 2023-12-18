package input

func ParseLines[T any](in *Input, parseFn func(string) T) []T {
	var items []T
	for in.Scan() {
		items = append(items, parseFn(in.Text()))
	}
	return items
}

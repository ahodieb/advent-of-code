package ints

func FromSpaceSeperated(s string) []int {
	var values []int
	n := 0
	in := false

	for _, r := range s {
		if '0' <= r && r <= '9' {
			if !in {
				in = true
				n = int(r - '0')
			} else {
				n = n*10 + (int(r - '0'))
			}
		} else {
			if in {
				values = append(values, n)
				in = false
				n = 0
			}
		}
	}

	if in {
		values = append(values, n)
	}

	return values
}

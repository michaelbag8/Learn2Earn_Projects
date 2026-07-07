package piscine

func Abort(a, b, c, d, e int) int {
	values := []int{a, b, c, d, e}
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values[2]
}

package piscine

func Map(f func(int) bool, a []int) []bool {
	boolslice := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			boolslice[i] = true
		} else {
			boolslice[i] = false
		}
	}
	return boolslice
}

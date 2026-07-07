package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || power > 20 {
		return 0
	}
	if nb == 0 && nb == 1 && power == 0 {
		return 1
	}

	result := 1

	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}

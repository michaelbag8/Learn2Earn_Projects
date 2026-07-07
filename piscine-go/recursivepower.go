package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if nb == 0 && power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	if nb == 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return nb * RecursivePower(nb, power-1)
}

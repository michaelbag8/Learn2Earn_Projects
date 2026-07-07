package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	if nb > 20 {
		return 0
	}
	factorial := 1
	for i := 2; i <= nb; i++ {
		factorial *= i
	}
	return factorial
}

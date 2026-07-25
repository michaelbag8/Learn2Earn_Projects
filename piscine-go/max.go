package piscine

func Max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number 
		}
	}
	return max
}

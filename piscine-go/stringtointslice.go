package piscine

func StringToIntSlice(str string) []int {
	var result []int
	for _, r := range str {
		con := int(r)
		result = append(result, con)
	}
	return result
}

package piscine

func Join(strs []string, sep string) string {
	var result string
	for i, r := range strs {
		result += r
		if i != len(strs)-1 {
			result += sep
		}
	}
	return result
}

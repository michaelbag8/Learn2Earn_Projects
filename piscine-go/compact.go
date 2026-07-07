package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			s[count] = s[i]
			count++
		}
	}
	*ptr = s[:count]
	return count
}

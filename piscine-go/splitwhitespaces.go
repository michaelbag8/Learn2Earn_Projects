package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	currentWord := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == ' ' || char == '\t' || char == '\n' {
			if currentWord != "" {
				result = append(result, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		result = append(result, currentWord)
	}
	return result
}

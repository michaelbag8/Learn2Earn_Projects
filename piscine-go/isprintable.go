package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if checks(runes[i]) == true {
			return false
		}
	}
	return true
}

func checks(x rune) bool {
	if x >= 0 && x <= 31 {
		return true
	}
	return false
}

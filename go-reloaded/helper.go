package main

import (
	"strconv"
	"strings"
)

func hexToDec(word string) string {
	val, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(val, 10)
}

func binToDec(word string) string {
	val, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(val, 10)
}

func applyCase(result []string, mode string, command string) []string {
	count := 1
	if index := strings.Index(command, ","); index != -1 {
		if c, err := strconv.Atoi(strings.Trim(command[index+1:], ")")); err == nil {
			count = c
		}
	}
	for i := len(result) - count; i < len(result); i++ {
		if i < 0 {
			continue
		}
		w := result[i]
		if mode == "up" {
			result[i] = strings.ToUpper(w)
		} else if mode == "low" {
			result[i] = strings.ToLower(w)
		} else if mode == "cap" && len(w) > 0 {
			result[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}
	return result
}

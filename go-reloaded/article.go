package main

import (
	"strings"
)

func fixArticles(text string) string {
	words := strings.Fields(text)
	vowels := "aeiouhAEIOUH"
	for i := 0; i < len(words)-1; i++ {
		isArticle := words[i] == "a" || words[i] == "A"
		nextStartsWithVowel := strings.ContainsRune(vowels, rune(words[i+1][0]))
		if isArticle && nextStartsWithVowel {
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}

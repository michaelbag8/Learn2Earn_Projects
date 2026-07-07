package main

import (
	"strings"
)

func multiPunctuation(text string) string {
	replacements := map[string]string{
		". . .": "...",
		"! !":   "!!",
		"? ?":   "??",
		"! ?":   "!?",
		"? !":   "?!",
	}

	for old, newVal := range replacements {
		text = strings.ReplaceAll(text, old, newVal)
	}

	return text
}
func fixPunctuation(text string) string {
	text = multiPunctuation(text)

	punctuations := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range punctuations {
		text = strings.ReplaceAll(text, " "+p, p)
		text = strings.ReplaceAll(text, p+"  ", p+" ")
	}
	text = strings.ReplaceAll(text, ",", ", ")
	text = strings.ReplaceAll(text, ":", ": ")

	for strings.Contains(text, "  ") {
		text = strings.ReplaceAll(text, "  ", " ")
	}

	return strings.TrimSpace(text)
}

func fixQuotes(text string) string {
	text = strings.ReplaceAll(text, "' ", "'")
	//text = strings.ReplaceAll(text, " '", "'")

	// fix cases where punctuation comes after the quote
	text = strings.ReplaceAll(text, " '.", "'.")
	text = strings.ReplaceAll(text, " ',", "',")
	text = strings.ReplaceAll(text, " '!", "'!")
	text = strings.ReplaceAll(text, " '?", "'?")
	text = strings.ReplaceAll(text, ". '", ".'")
	//text = strings.ReplaceAll(text, ", '", ",'")

	return text
}

package main

import (
	"strings"
)

func renderLines(segment string, blockMaps map[rune][]string) []string {
	var output []string

	for row := 0; row < 8; row++ {
		var rawRow strings.Builder
		for _, ch := range segment {
			rawRow.WriteString(blockMaps[ch][row])
		}

		output = append(output, rawRow.String())
	}
	return output
}

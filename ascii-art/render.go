package main

import "strings"


func renderLines(segment string, blockMaps map[rune][]string) []string {
	var output []string

	for row := range 8{
		var result strings.Builder
		for _, ch := range segment {
			result.WriteString(blockMaps[ch][row])
		}

		output = append(output, result.String())
	}
	return output
}

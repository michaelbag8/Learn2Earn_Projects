package main

import (
	"strings"
)

func Generate(input string, banner map[rune][]string) string {

	var output strings.Builder
	segments := SplitInput(input)

	for i, segment := range segments {
		if segment == "" {
			if i < len(segments)-1 {
				output.WriteString("\n")
			}

			continue
		}

		rows := renderLines(segment, banner)
		for _, row := range rows {
			output.WriteString(row)
			output.WriteString("\n")
		}
		if i < len(segments)-1 {
			output.WriteString("\n")
		}

	}
	return output.String()
}

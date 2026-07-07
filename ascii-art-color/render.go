package main

import (
    "strings"
   
)


func renderLines(segment string, blockMaps map[rune][]string, colorCode string, letters string) []string {
    var output []string

        for row := 0; row < 8; row++ {
            var rawRow strings.Builder
            for _, ch := range segment {
                
                shouldColor := colorCode != "" && (letters == "" || strings.ContainsRune(letters, ch))
                if shouldColor {
                    rawRow.WriteString(colorCode)
                    rawRow.WriteString(blockMaps[ch][row])
                    rawRow.WriteString("\033[0m")
                } else {
                    rawRow.WriteString(blockMaps[ch][row])
                }
                
            }
            output = append(output, rawRow.String())
        }
        return output
    }
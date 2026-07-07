package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	color := flag.String("color", "", "colors: red, green, yellow, cyan ...")

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: go run . --color=color [LETTERS] TEXT [BANNER]")
		os.Exit(1)
	}

	var letters string
	var input string
	fontFile := "standard.txt"

	// normal case
	if len(args) == 1 {
		input = args[0]
		letters = "" 
	}

	// with substring
	if len(args) >= 2 {
		letters = args[0]
		input = args[1]
	}

	// with banner file
	if len(args) >= 3 {
		fontFile = args[2]+".txt"
	}

	// color handling
	colorCode := ""

	if *color != "" {
		code, ok := colorMap[*color]
		if !ok {
			fmt.Fprintf(os.Stderr, "unknown color: %s\n", *color)
			os.Exit(1)
		}
		colorCode = code
	}

	data, err := LoadBanner(fontFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading banner: %v\n", err)
		os.Exit(1)
	}

	content := Generate(input, data, colorCode, letters)
	fmt.Print(content)
}
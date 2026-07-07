package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	output := flag.String("output", "", "file to write output to")

	flag.Parse()

	args := flag.Args()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: go run . [--output=file] <text> [banner]\n")
	}

	input := args[0]

	fontFile := "standard.txt"
	if len(args) > 1 {
		fontFile = args[1]+".txt"
	}

	data, err := LoadBanner(fontFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing file: %v\n", err)
		return
	}

	
	content := Generate(input, data)

	if *output != "" {
		err := os.WriteFile(*output, []byte(content), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error : writing file %v\n", err)
			return
		}
		fmt.Printf("Output written to %s\n", *output)
	} else {
		fmt.Print(content)
	}

}

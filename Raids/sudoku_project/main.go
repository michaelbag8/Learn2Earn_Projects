package main

import (
	"fmt"

	"os"
)

func main() {
	args := os.Args[1:]

	board, ok := ParseBoard(args)
	if !ok {
		fmt.Println("Error")
		return
	}

	solved, ok := SolveUnique(board)
	if !ok {
		fmt.Println("Error")
		return
	}

	PrintBoard(solved)
}

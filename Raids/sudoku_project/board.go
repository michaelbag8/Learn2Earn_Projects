package main

import "fmt"

type Board [9][9]int

// Convert command line args to a board
func ParseBoard(args []string) (Board, bool) {
	if len(args) != 9 {
		return Board{}, false
	}

	var b Board

	for i := 0; i < 9; i++ {
		if len(args[i]) != 9 {
			return Board{}, false
		}
		for j := 0; j < 9; j++ {
			ch := args[i][j]
			if ch == '.' {
				b[i][j] = 0
			} else if ch >= '1' && ch <= '9' {
				b[i][j] = int(ch - '0')
			} else {
				return Board{}, false
			}
		}
	}

	return b, true
}

// Print solved board in required format
func PrintBoard(b Board) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(b[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

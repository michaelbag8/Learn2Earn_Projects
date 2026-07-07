package piscine

import "fmt"

func QuadD(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {

			// checking through the top-left
			if i == 0 && j == 0 {
				fmt.Print("A")

				// checking through the top-right
			} else if i == 0 && j == x-1 {
				fmt.Print("C")

				// checking through the bottom-left
			} else if i == y-1 && j == 0 {
				fmt.Print("A")

				// checking through the bottom-right
			} else if i == y-1 && j == x-1 {
				fmt.Print("C")

				// checking through the top or bottom border
			} else if i == 0 || i == y-1 {
				fmt.Print("B")

				// checking through the left or right border
			} else if j == 0 || j == x-1 {
				fmt.Print("B")

			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

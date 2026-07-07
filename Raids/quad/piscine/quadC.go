package piscine

import "fmt"

func QuadC(x, y int) {
	if x <= 0 || y <= 0 { // to return nothing if value is less than 0
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {

			// top-left
			if i == 0 && j == 0 {
				fmt.Print("A")

				// top-right
			} else if i == 0 && j == x-1 {
				fmt.Print("A")

				// bottom-left
			} else if i == y-1 && j == 0 {
				fmt.Print("C")

				// bottom-right
			} else if i == y-1 && j == x-1 {
				fmt.Print("C")

				// top or bottom border
			} else if i == 0 || i == y-1 {
				fmt.Print("B")

				// left or right border
			} else if j == 0 || j == x-1 {
				fmt.Print("B")

			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

package piscine

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 { // to return nothing if value is less than 0
		return
	}

	for i := 0; i < y; i++ { // going through height
		for j := 0; j < x; j++ { // going throug width
			if (i == 0 && j == 0) || (i == 0 && j == x-1) || // checking for the corners
				(i == y-1 && j == 0) || (i == y-1 && j == x-1) {
				fmt.Print("o")
			} else if i == 0 || i == y-1 { /// checking for top and bottom side
				fmt.Print("-")
			} else if j == 0 || j == x-1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

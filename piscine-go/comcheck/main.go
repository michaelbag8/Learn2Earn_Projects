package piscine

import (
	"fmt"
	"os"
)

func main() {
	for v := 0; v < len(os.Args); v++ {
		arg := os.Args[v]
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}

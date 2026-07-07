package main

import (
	"os"
	"piscine/piscine"
	"strconv"
)

func main() {
	if len(os.Args) != 4 { // If the number of arguments is not 4, exit the program
		return
	}
	quadType := os.Args[1]                  // store the first argument as quadTpye
	width, err1 := strconv.Atoi(os.Args[2]) // convert second argument from str to int
	height, err2 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil { // If any of the conversions didn't, the function should exit
		return
	}
	switch quadType {
	case "A": // If quadType is "A", the QuadA function from the piscine package should be call
		piscine.QuadA(width, height)

	case "B":
		piscine.QuadB(width, height)

	case "C":
		piscine.QuadC(width, height)

	case "D":
		piscine.QuadD(width, height)

	case "E":
		piscine.QuadE(width, height)
	}
}

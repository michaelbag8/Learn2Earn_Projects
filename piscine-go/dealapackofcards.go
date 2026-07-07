package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for player := 0; player <= 4; player++ {
		fmt.Printf("Player %d: ", player)
		start := player * 4
		end := start + 4

		deckLength := 0
		for range deck {
			deckLength++
		}

		for i := start; i < end && i < deckLength; i++ {
			if i == end-1 {
				fmt.Printf("%d", deck[i])
			} else {
				fmt.Printf("%d, ", deck[i])
			}
		}

		z01.PrintRune('\n')
	}
}

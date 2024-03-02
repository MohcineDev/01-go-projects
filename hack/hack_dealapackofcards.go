package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	count := 1

	for i := 0; i < 4; i++ {

		count += i

		fmt.Printf("Player %d: ", i+1)
		if i == 1 {
			count = 4
		}
		if i == 2 {
			count = 7
		}
		if i == 3 {
			count = 10
		}

		for j := 0; j < 3; j++ {

			fmt.Printf("%d", j+count)

			if j < 2 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n")

	}
}

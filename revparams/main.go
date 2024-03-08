package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		mm := os.Args[i]
		for k := 0; k < len(mm); k++ {
			z01.PrintRune(rune(mm[k]))
		}
		z01.PrintRune(rune('\n'))
	}
}

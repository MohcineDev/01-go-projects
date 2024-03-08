package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	result := os.Args
	j := []rune(result[0])
	for i := 2; i < len(j); i++ {
		z01.PrintRune(j[i])
	}
	z01.PrintRune('\n')
}

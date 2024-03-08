package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for j, v := range os.Args {
		if j > 0 {
			for _, elem := range v {
				z01.PrintRune(elem)
			}
			z01.PrintRune('\n')
		}
	}
}

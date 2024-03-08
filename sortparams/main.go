package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myslice := []string{}
	for i := 0; i < len(os.Args); i++ {
		if i > 0 {
			myslice = append(myslice, os.Args[i])
		}
	}

	for i := 0; i < len(myslice); i++ {
		for j := i + 1; j < len(myslice); j++ {
			if myslice[i] > myslice[j] {
				temp := myslice[i]
				myslice[i] = myslice[j]
				myslice[j] = temp
			}
		}
	}
	for i := 0; i < len(myslice); i++ {
		for j := 0; j < len(myslice[i]); j++ {
			z01.PrintRune(rune(myslice[i][j]))
		}
		z01.PrintRune('\n')
	}
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myArgs := os.Args

	toAdd := 96
	isFlaged := 1
	if len(os.Args) <= 1 {
		return
	}
	if os.Args[1] == "--upper" {
		toAdd = 64
		isFlaged = 2
	}

	for i := isFlaged; i < len(myArgs); i++ {
		// chack if the number is within the alphabet range 0 - 26
		if convertStrToInt(myArgs[i]) > 0 && convertStrToInt(myArgs[i]) <= 26 {
			total := toAdd + convertStrToInt(myArgs[i])
			z01.PrintRune(rune(total))

		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

///convert string to int

func convertStrToInt(s string) int {
	result := 0
	for _, v := range s {
		result *= 10
		if v >= '0' && v <= '9' {
			result += int(v - 48)
		}
	}
	return result
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

type boolean struct {
	value int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	var yes boolean
	var no boolean
	yes.value = 1
	no.value = 0

	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func even(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}

func main() {
	lengthOfArg := len(os.Args[1:])

	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	if isEven(lengthOfArg).value == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

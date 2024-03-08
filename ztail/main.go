package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myArgs := os.Args[1:]
	if len(myArgs) >= 1 {
		for i := 2; i < len(myArgs); i++ {

			content, err := os.ReadFile(myArgs[i])
			if err != nil {
			} else {
				option := convertStringToInt(myArgs[1])
				// zyx
				for i := 0; i < len(content); i++ {
					lastChar := content[len(content)-1-i]
					if i == option {
						break
					}
					if lastChar != '\n' {
						z01.PrintRune(rune(lastChar))
					}
				}
				z01.PrintRune('\n')

			}
		}
	}
}

func convertStringToInt(a string) int {
	result := 0
	for _, v := range a {
		result *= 10
		if v >= '0' && v <= '9' {
			result += int(v - 48)
		}
	}
	return result
}

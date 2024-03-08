package main

import (
	//	"fmt"
	"os"
)

func main() {
	myArgs := os.Args[1:]
	result := 5
	if len(myArgs) == 3 {
		if myArgs[1] == "+" {
			result = stringToInt(myArgs[0]) + stringToInt(myArgs[2])
		} else if myArgs[1] == "-" {
			result = stringToInt(myArgs[0]) - stringToInt(myArgs[2])
		} else if myArgs[1] == "/" {
			result = stringToInt(myArgs[0]) / stringToInt(myArgs[2])
		} else if myArgs[1] == "*" {
			result = stringToInt(myArgs[0]) * stringToInt(myArgs[2])
		} else if myArgs[1] == "%" {
			result = stringToInt(myArgs[0]) % stringToInt(myArgs[2])
		}
	}
	myArgs = append(myArgs, result)
}

func stringToInt(s string) int {
	R := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			R *= 10
			R += int(v - '0')
		}
	}
	return R
}

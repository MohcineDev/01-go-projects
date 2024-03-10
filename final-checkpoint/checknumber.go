package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

func CheckNumber(arg string) bool {
 
	result := false

	if len(arg) >= 1 {
		
		for i := 0; i < len(arg); i++ {
			if arg[i] >= '0' &&  arg[i] <= '9' {
				result = true
				break
			}
		}
		
	}  
	return result
}
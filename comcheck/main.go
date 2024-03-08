package main

import (
	"fmt"
	"os"
)

func main() {
	myArgs := os.Args[1:]
	toMatch := []string{"01", "galaxy", "galaxy 01"}

	for i := 0; i < len(toMatch); i++ {
		for j := 0; j < len(myArgs); j++ {
			if toMatch[i] == myArgs[j] {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}

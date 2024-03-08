package main

import (
	"fmt"
	"os"
)

func main() {
	myArgs := os.Args[1:]
	if len(myArgs) == 0 {
		fmt.Println("File name missing")
		return
	} else if len(myArgs) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	body, _ := os.ReadFile(myArgs[0])

	fmt.Print(string(body))
}

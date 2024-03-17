package main

import (
	"fmt"
	"os"
)

func main() {
	brackets()
}

func brackets() {
	myArgs := os.Args[1:]
 
	opening := []string {"(", "{", "["}
	closing := []string {")", "}", "]"}
	argsOpening := []int {}
	argsClosing := []int {}
///([)]

	for i := 0; i < len(myArgs); i++ {
		for j := 0; j < len(myArgs[i]); j++ {
			for k := 0; k < len(opening); k++ {
				if string(myArgs[i][j]) == opening[k]{
					fmt.Println(string(myArgs[i][j]))

					argsOpening = append(argsOpening, int(myArgs[i][j]) )	
				}
			}
			for k := 0; k < len(closing); k++ {
				if string(myArgs[i][j]) == closing[k]{
					fmt.Println(string(myArgs[i][j]))

					argsClosing = append(argsClosing, int(myArgs[i][j]) )	
				}
			}	
		}
	}
	fmt.Println(argsOpening)
	fmt.Println(argsClosing)
	///check if elems in argsOpening have correspondent elems in closing
	correct := true
	O :=len(argsOpening)
	if len(argsOpening) > 0{
		for i := 0; i < len(argsOpening); i++ {
			// for j := len(argsClosing)-1; j >=0 ; j++ {
				if argsOpening[i] == 40 {
					if argsOpening[i] != argsClosing[O-1]-1{
						correct = false
					}
				} else if argsOpening[i] == 91 || argsOpening[i] == 123 {
					if argsOpening[i] == argsClosing[O-1]-2{
						correct = false
					}
				}	
			// }	
		}
	}
	if correct {
		fmt.Println("OK")
		
	} else{
		fmt.Println("Error")

	}

}
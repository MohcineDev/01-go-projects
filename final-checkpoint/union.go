package main

import (
	"fmt"
	"os"
)

func main() { 
	Union()
}
 
func Union() {
	
	myArgs := os.Args[1:]
	result :=""

	if len(myArgs) != 2 {
		fmt.Print("\n")
	} else{

		result0 := addChars(myArgs[0], result)
		result +=  addChars(myArgs[1], result0)
		   		
		fmt.Print(result + "\n")
	}
}

func addChars(str , result string) string{
	found := false
	result += string(str[0])

	for i := 0; i < len(str); i++ {
		found = false
		
		for j := 0; j < len(result); j++ {
			if string(str[i]) == string(result[j]) {
				found = true
				break
			}
		}
		if !found {
			result += string(str[i])
		}
	}

	return result
}
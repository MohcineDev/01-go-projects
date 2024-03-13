package main

import (
	"fmt"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(str string) string {

	newString := []string{}
	result := ""
	
	if str == "" {
		return "Invalid Output\n"
	} else {

		word:=""
		
		for i := 0; i < len(str); i++ {
			if string(str[i]) != " " {
				word += string(str[i])
			} 
			 if word != "" && string(str[i]) == " " || i == len(str)-1{
				newString = append(newString, word)
				if i != len(str)-1{
					newString = append(newString, " ")
				}
				word = ""
			}
		}
		///flip
		for i := len(newString)-1; i >=0  ; i-- {
			result += newString[i]
		}
	}
	return result+"\n"
}
package main

import (
	"fmt"
)

func main() {
	 
	s := "HelloxDhowxDarexDyou?"
	fmt.Printf("%#v\n",Split(s, "xD"))
}

func Split(s, sep string) []string {
	result :=[] string{}
	found := true
	start := 0
	
	for i := 0; i < len(s); i++ {

		if s[i] == sep[0] {
			found = true
			for j := 1; j < len(sep); j++ {
				if s[i+j] != sep[j] {
					found = false
					break
				}
			}
			if found  {
				result = append(result , s[start: i])
				start = i + len(sep)
			}
		}
		if i == len(s)-1 && len(s) > start {
			result = append(result , s[start:])
		}
 
	}

	return result
}
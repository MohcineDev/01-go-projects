package main

import "fmt"

func main(){
	fmt.Println(	ItoaBase(4645468, 11))
}


func ItoaBase(value, base int) string {
	theBase := "0123456789ABCDEF"
	sign :=""
	result:=""
	
	if value < 0 {
		sign = "-"
		value *= -1
	}

	for value > 0 {
		result =  string(theBase[value % base]) + result
//		fmt.Println("value %= 10 : ", value % base, string(theBase[value % base]))
		value /= base

	}
	fmt.Println("itoa : ", sign + result)
	return sign + result
}

///hint
//////https://www.rapidtables.com/convert/number/base-converter.html
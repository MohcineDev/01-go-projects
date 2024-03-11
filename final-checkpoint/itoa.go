package main

import (
	"fmt"
)

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {

	result :=""
	sign :=""

	if n == 0 {
		result = "0"
	}
	if n < 0 {
		sign = "-"
		n = -n
	}
	for n > 0{
		a := n%10
		n /= 10
		result = string(a + '0') + result
	}
	return sign + result
}
package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	myInt := []int{}

	// while the number is > 9

	for n > 9 {
		myInt = append(myInt, n%10)
		n = n / 10
	}

	if n <= 9 {
		myInt = append(myInt, n)
	}

	/// swap the values in ascending order
	for i := 0; i < len(myInt); i++ {
		for j := i + 1; j < len(myInt); j++ {
			if myInt[i] > myInt[j] {
				temp := myInt[i]
				myInt[i] = myInt[j]
				myInt[j] = temp
			}
		}
	}

	///print the result after the swap
	for _, v := range myInt {
		z01.PrintRune(rune(v + '0'))
	}
}

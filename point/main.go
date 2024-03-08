package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

const a = "x = 42, y = 21"

func main() {
	points := &point{}

	// for i := 0; i < 14; i++ {
	// 	z01.PrintRune(rune(a[i]))
	// }
	for _, v := range a {
		z01.PrintRune(v)
	}

	z01.PrintRune('\n')

	setPoint(points)
}

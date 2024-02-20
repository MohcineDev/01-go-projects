package main

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x < 0 || y < 0 {
		return
	} else {

		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				if j == 0 && i == 0 || j == y-1 && i == 0 || j == 0 && i == x-1 || j == y-1 && i == x-1 {
					z01.PrintRune('o')
				} else if j == 0 && i > 0 && i < x-1 || j == y-1 && i > 0 && i < x-1 {
					z01.PrintRune('-')
				} else if i == 0 && j > 0 && j < y-1 || i == x-1 && j > 0 && j < y-1 {
					z01.PrintRune('|')
				} else if i > 0 && j > 0 && i < x && j < y {
					z01.PrintRune(' ')
				}
			}

			z01.PrintRune('\n')
		}
		z01.PrintRune('\n')
	}
}

/*
for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 || i==y {
					if j==1 || j==x{
						z01.PrintRune('o')
					}else{
						z01.PrintRune('-')
					}
				} else {
					if j ==1 || j==x{
						z01.PrintRune('|')
					}else{
						z01.PrintRune(' ')
					}
				}
			}

*/

/*
   	  i=0  i=1  i=2  i=3   X
   j=0 o    -     -    o
   j=1 |               |
   j=2 |               |
   j=3 o    -     -    o
   Y
*/

func QuadB(x, y int) {
	if x < 0 || y < 0 {
		return
	} else {
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				if j == 0 && i == 0 || j == y-1 && i == x-1 && i > 0 && j > 0 {
					z01.PrintRune('/')
				} else if j == 0 && i == x-1 || j == y-1 && i == 0 {
					z01.PrintRune('\\')
				} else if j == 0 && i > 0 && i < x-1 || j == y-1 && i > 0 && i < x-1 || i == 0 && j > 0 && j < y-1 || i == x-1 && j > 0 && j < y-1 {
					z01.PrintRune('*')
				} else if j > 0 && j < y-1 && i > 0 && i < x-1 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
		z01.PrintRune('\n')
	}
}

func QuadC(x, y int) {
	if x < 0 || y < 0 {
		return
	} else {
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				if j == 0 && i == 0 || j == 0 && i == x-1 {
					z01.PrintRune('A')
				} else if j == y-1 && i == 0 || j == y-1 && i == x-1 && i > 0 && j > 0 {
					z01.PrintRune('C')
				} else if j == 0 && i > 0 && i < x-1 || j == y-1 && i > 0 && i < x-1 || i == 0 && j > 0 && j < y-1 || i == x-1 && j > 0 && j < y-1 {
					z01.PrintRune('B')
				} else if j > 0 && j < y-1 && i > 0 && i < x-1 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
		z01.PrintRune('\n')
	}
}

func QuadD(x, y int) {
	if x < 0 || y < 0 {
		return
	} else {
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				if j == 0 && i == 0 || j == y-1 && i == 0 {
					z01.PrintRune('A')
				} else if j == 0 && i == x-1 || j == y-1 && i == x-1 && i > 0 && j > 0 {
					z01.PrintRune('C')
				} else if j == 0 && i > 0 && i < x-1 || j == y-1 && i > 0 && i < x-1 || i == 0 && j > 0 && j < y-1 || i == x-1 && j > 0 && j < y-1 {
					z01.PrintRune('B')
				} else if j > 0 && j < y-1 && i > 0 && i < x-1 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
		z01.PrintRune('\n')
	}
}

func QuadE(x, y int) {
	if x < 0 || y < 0 {
		return
	} else {
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				if j == 0 && i == 0 || j == y-1 && i == x-1 && i > 0 && j > 0 {
					z01.PrintRune('A')
				} else if j == y-1 && i == 0 || j == 0 && i == x-1 {
					z01.PrintRune('C')
				} else if j == 0 && i > 0 && i < x-1 || j == y-1 && i > 0 && i < x-1 || i == 0 && j > 0 && j < y-1 || i == x-1 && j > 0 && j < y-1 {
					z01.PrintRune('B')
				} else if j > 0 && j < y-1 && i > 0 && i < x-1 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
		z01.PrintRune('\n')
	}
}

func main() {
	// QuadA(5, 3)

	// QuadA(5, 3)
	// QuadA(5, 1)
	// QuadA(1, 1)
	// QuadA(1, 5)
	//---
	// QuadB(5, 3)
	// QuadB(5, 1)
	// QuadB(1, 1)
	// QuadB(1, 5)
	//---
	// QuadC(5,3)
	// QuadC(5,1)
	// QuadC(1,1)
	// QuadC(1,5)
	//---
	// QuadD(5, 3)
	// QuadD(5, 1)
	// QuadD(1, 1)
	// QuadD(1, 5)
	//---
	QuadE(5, 3)
	QuadE(5, 1)
	QuadE(1, 1)
	QuadE(1, 5)
}

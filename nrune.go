package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	} else {
		for i, elem := range s {
			if i+1 == n {
				return elem
			}
		}
	}
	return 0
}

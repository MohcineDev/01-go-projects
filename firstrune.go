package piscine

func FirstRune(s string) rune {
	for _, elem := range s {
		return elem
	}

	return ' '
}

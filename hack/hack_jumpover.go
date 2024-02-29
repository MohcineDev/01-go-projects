package piscine

func JumpOver(str string) string {
	toPrint := ""
	newLine := "\n"
	if len(str) <= 2 {
		return newLine
	} else {
		for i := 1; i <= len(str); i++ {
			if i%3 == 0 {
				toPrint += string(str[i-1])
				// z01.PrintRune(rune(str[i-1]))
			}
		}
	}

	return toPrint + "\n"
}

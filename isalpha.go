package piscine

func IsAlpha(s string) bool {
	nbr := 0

	for j := 0; j < len(s); j++ {
		for i := 'A'; i <= 'Z'; i++ {
			if rune(s[j]) == i {
				nbr++
				break
			}
		}
		for i := 'a'; i <= 'z'; i++ {
			if rune(s[j]) == i {
				nbr++
				break
			}
		}

		if '0' <= s[j] && '9' >= s[j] {
			nbr++
		}
	}

	if nbr == len(s) {
		return true
	}
	return false
}

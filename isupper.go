package piscine

func IsUpper(s string) bool {
	nbr := 0

	for j := 0; j < len(s); j++ {
		for i := 'A'; i <= 'Z'; i++ {
			if rune(s[j]) == i {
				nbr++
				break
			}
		}
	}
	if nbr == len(s) {
		return true
	}
	return false
}

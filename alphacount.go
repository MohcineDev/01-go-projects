package piscine

func AlphaCount(s string) int {
	totalNbr := 0
	found := true
	for i := 0; i < len(s); i++ {
		found = true

		for j := 'a'; j < 'z'; j++ {
			if rune(s[i]) == j {
				totalNbr++
				found = !found
			}
		}
		if found {
			for j := 'A'; j < 'Z'; j++ {
				if rune(s[i]) == j {
					totalNbr++
				}
			}
		}
	}
	return totalNbr
}

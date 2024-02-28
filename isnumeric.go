package piscine

func IsNumeric(s string) bool {
	result := true

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			result = false
			break
		}
	}

	return result
}

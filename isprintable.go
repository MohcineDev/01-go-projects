package piscine

func IsPrintable(s string) bool {
	result := true
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\b' || s[i] == '\f' {
			result = false
			break
		}
	}
	return result
}

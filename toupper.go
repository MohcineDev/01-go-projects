package piscine

func ToUpper(s string) string {
	var result string

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			alphaNum := v - 32
			result += string(alphaNum)
		} else {
			result += string(v)
		}
	}

	return result
}

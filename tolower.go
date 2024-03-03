package piscine

func ToLower(s string) string {
	var result string

	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			num := v + 32
			result += string(num)
		} else {
			result += string(v)
		}
	}

	return result
}

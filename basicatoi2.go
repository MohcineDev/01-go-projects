package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			result *= 10
			result += int(v - 48)
		} else {
			return 0
		}
	}
	return result
}

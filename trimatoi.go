package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	firstSign := true

	for _, v := range s {
		if v == '-' && firstSign {
			sign = -1
		} else if v >= '0' && v <= '9' {
			firstSign = false
			result *= 10
			result += int(v - '0')
		}
	}

	return result * sign
}

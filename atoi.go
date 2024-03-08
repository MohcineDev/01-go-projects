package piscine

func Atoi(s string) int {
	result := 0
	sign := 1

	if len(s) == 0 || s == "" {
		return 0
	}
	i := 0
	if s[i] == '-' {
		sign = -1
		i++
	}
	if s[i] == '+' {
		sign = 1
		i++
	}

	for ; i < len(s); i++ {

		if s[i] >= '0' && s[i] <= '9' {
			result *= 10
			result += int(s[i] - 48)

		} else {
			return 0
		}
	}
	return result * sign
}

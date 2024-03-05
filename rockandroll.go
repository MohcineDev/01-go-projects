package piscine

func RockAndRoll(n int) string {
	result := ""

	if n < 0 {
		result = "error: number is negative\n"
	} else if n%2 == 0 && n%3 == 0 {
		result = "rock and roll\n"
	} else if n%2 == 0 {
		result = "rock\n"
	} else if n%3 == 0 {
		result = "roll\n"
	} else {
		result = "error: non divisible\n"
	}

	return result
}

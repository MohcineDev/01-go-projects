package piscine

func Compare(a, b string) int {
	result := 0

	if a == b {
		result = 0
	} else if a < b {
		result = -1
	} else if a > b {
		result = 1
	}
	return result
}

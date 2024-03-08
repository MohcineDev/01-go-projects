package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 1 {
		for i := 1; i < nb+1; i++ {
			result *= i
			// if result < 0 {
			// 	return 0
			// }
		}
	} else if nb == 0 {
		return 1
	} else if nb == 1 {
		return 1
	} else {
		return 0
	}
	return result
}

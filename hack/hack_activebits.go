package piscine

func ActiveBits(n int) int {
	// if n == 0 {
	// 	return 0
	// }
	// return ActiveBits(n/2) + n%2

	count := 0
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}

	return count
}

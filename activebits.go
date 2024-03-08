package piscine

func ActiveBits(n int) int {
	if n == 0 {
		return 0
	}
	return ActiveBits(n/2) + n%2
}

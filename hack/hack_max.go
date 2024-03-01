package piscine

func Max(a []int) int {
	big := 0
	if len(a) < 1 {
		return 0
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] > big {
				big = a[i]
			}
		}
	}

	return big
}

package piscine

func Unmatch(a []int) int {
	count := 0

	for _, v := range a {
		for _, e := range a {
			if v == e {
				count++
			}
		}
		if count == 1 || count%2 == 1 {
			return v
		}
		count = 0
	}
	return -1
}

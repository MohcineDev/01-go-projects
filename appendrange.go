package piscine

func AppendRange(min, max int) []int {
	s := []int{}

	if min >= max {
		return []int(nil)
	}
	for i := min; i < max; i++ {
		s = append(s, i)
	}

	return s
}

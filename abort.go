package piscine

func Abort(a, b, c, d, e int) int {
	nbrs := []int{}

	nbrs = append(nbrs, a, b, c, d, e)
	for i := 0; i < len(nbrs); i++ {
		for j := 0; j < len(nbrs); j++ {
			if nbrs[i] > nbrs[j] {
				temp := nbrs[i]
				nbrs[i] = nbrs[j]
				nbrs[j] = temp
			}
		}
	}

	return nbrs[len(nbrs)/2]
}

package piscine

func Unmatch(a []int) int {
	count := 0

	for _, v := range a {
		for _, e := range a {
			if v == e {
				count++
			}
		}
		if count == 1 || count%2 == 1 { // 4 4 4
			return v
		}
		count = 0

	}
	/*	matches := []int{}
		for i := 0; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				if a[i] == a[j] {
					// a[i] = -1
					// a[j] = -1
					fmt.Println("a[i] : ", a[i])
					matches = append(matches, a[i])
				}
			}
		}
		fmt.Println("a : ", a)
		fmt.Println("slice : ", matches)

		for i := 0; i < len(a); i++ {
			found := false
			for j := 0; j < len(matches); j++ {
				if a[i] == matches[j] {
					found = true
				}
			}
			if !found {
				return a[i]
			}
		}

	*/
	return -1
}

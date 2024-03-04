package piscine

func Split(s, sep string) []string {
	result := []string{}
	indexes := []int{}

	for i := 0; i < len(s); i++ {
		found := true
		if s[i] == sep[0] {
			for j := 0; j < len(sep); j++ {
				if sep[j] != s[i+j] {
					found = false
					break
				}
			}
			// if we found the sep
			if found {
				indexes = append(indexes, i)
			}
		}
	}
	start := 0
	for i := 0; i < len(indexes); i++ {
		var toAdd string

		toAdd = s[start:indexes[i]]

		start = indexes[i] + len(sep)
		result = append(result, toAdd)
		///use len(sep) becuase the sep may change to more that 2 chars ex: AH, bb, !==!
		if start-len(sep) == indexes[len(indexes)-1] {
			toAdd = s[indexes[i]+len(sep):]
			result = append(result, toAdd)
		}

	}

	return result
}

package piscine

func Index(s string, toFind string) int {
	equal := false
	// fmt.Println(".")
	startIndex := -1
	if len(s) < len(toFind) || len(s) < 1 {
		return -1
	} else if len(toFind) == 0 {
		return 0
	} else {
		for i := 0; i < len(s); i++ {
			count := 1
			if s[i] == toFind[0] {
				equal = true
				//				fmt.Println("i : ", i)

				startIndex = i // 0
				for j := 1; j < len(toFind); j++ {
					if i+j < len(s)-1 {
						if s[i+j] == toFind[j] {
							count = count + 1
						} else {
							equal = false
						}
					}
				}
				if count == len(toFind) {
					return i
				}

			} else if equal == false && i == len(s)-1 {
				return -1
			}
		}
	}
	return startIndex
}

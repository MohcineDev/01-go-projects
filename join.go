package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i := 0; i < len(strs); i++ {
		strAndSep := ""
		if i < len(strs)-1 {
			strAndSep = strs[i] + sep
			result += strAndSep
		} else {
			result += strs[i]
		}
	}

	return result
}

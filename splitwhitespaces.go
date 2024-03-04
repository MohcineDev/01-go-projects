package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	var word string
	for i := 0; i < len(s); i++ {
		if s[i] != 9 && s[i] != 10 && string(s[i]) != " " {
			// fill word with chars
			word += string(s[i])
		}
		if string(s[i]) == " " || i == len(s)-1 || s[i] == 9 || s[i] == 10 {
			///append to slice
			if len(word) > 0 {
				result = append(result, word)
			}
			word = ""
		}
	}
	return result
}

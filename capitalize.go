package piscine

func Capitalize(s string) string {
	myStr := []rune(s)
	result := ""

	for i, v := range myStr {
		if checkIfAlphaNumeric(v) == "upper" {
			myStr[i] = myStr[i] + 32
			result += string(myStr[i])
		}
	}

	for i, v := range myStr {
		if i == 0 && checkIfAlphaNumeric(v) == "lower" {
			myStr[i] = myStr[i] - 32
			
		}
		if i != 0 && checkIfAlphaNumeric(v) == "lower" && checkIfAlphaNumeric(myStr[i-1]) == "notAlphaNum" {
			myStr[i] = myStr[i] - 32
		} else {
			result += string(myStr[i])
		}
	}

	return string(myStr)
}

func checkIfAlphaNumeric(char rune) string {
	cat := "notAlphaNum"

	if char >= 'A' && char <= 'Z' {
		cat = "upper"
	} else if char >= 'a' && char <= 'z' {
		cat = "lower"
	} else if char >= '0' && char <= '9' {
		cat = "num"
	}

	return cat
}

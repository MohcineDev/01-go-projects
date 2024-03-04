package piscine

func StrRev(s string) string {
	myreversedString := ""

	for i := len(s) - 1; i >= 0; i-- {
		myreversedString += string(s[i])
	}

	return myreversedString
}

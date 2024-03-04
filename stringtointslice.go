package piscine

func StringToIntSlice(str string) []int {
	mySlice := []int(nil)

	for _, v := range str {
		mySlice = append(mySlice, int(v))
	}

	return mySlice
}

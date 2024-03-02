package piscine

func Compact(ptr *[]string) int {
	// count := 0
	newSlice := []string{}

	for _, v := range *ptr {
		if v != "" {
			// count++
			newSlice = append(newSlice, v)
		}
	}
	*ptr = newSlice

	return len(newSlice)
}

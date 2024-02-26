package piscine

func ConcatParams(args []string) string {
	result := ""
	for index, v := range args {
		if index != len(args)-1 {
			result += v + "\n"
		} else {
			result += v
		}
	}
	return result
}

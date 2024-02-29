package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	grocery := make(map[string]int)
	items := []string{}
	item := ""
	count := 0

	// for i, v := range str {
	// 	if v != ' ' {
	// 		item += string(v)
	// 	}
	// 	if v == ' ' && i+1 != ' ' && v != 9 && item != "" || i == len(str)-1 {
	// 		items = append(items, item)
	// 		item = ""

	// 	}
	// }
	for i, v := range str {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			item += string(v)
		}
		if v == ' ' && i+1 != ' ' && v != 9 && item != "" || i == len(str)-1 {
			items = append(items, item)
			item = ""

		}
	}

	for i := 0; i < len(items); i++ {

		_, ok := grocery[items[i]]
		if ok {
			grocery[items[i]] = grocery[items[i]] + 1
		} else {
			grocery[items[i]] = 1
		}
		// }

		count++

	}
	return grocery
}

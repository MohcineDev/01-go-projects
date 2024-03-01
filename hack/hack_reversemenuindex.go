package piscine

func ReverseMenuIndex(menu []string) []string {
	newMenu := []string{}
	for i := len(menu) - 1; i >= 0; i-- {
		newMenu = append(newMenu, menu[i])
	}
	return newMenu
}

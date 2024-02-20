package piscine

func UltimateDivMod(a *int, b *int) {
	temp := *a / *b
	*b = *a % *b
	*a = temp
}

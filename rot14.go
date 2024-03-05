package piscine

func Rot14(s string) string {
	var R string
	for _, v := range s {
		R += string(print14(v))
	}
	return R
}

func print14(v rune) rune {
	if v >= 'A' && v < 'M' || v >= 'a' && v < 'm' {
		return v + 14
	}
	if v >= 'M' && v <= 'Z' || v >= 'm' && v <= 'z' {
		return v - 12
	}
	return v
}

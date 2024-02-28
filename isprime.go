package piscine

func IsPrime(nb int) bool {
	isPrimeBool := true
	if nb <= 1 {
		isPrimeBool = !isPrimeBool
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				isPrimeBool = false
				break
			}
		}
	}
	return isPrimeBool
}

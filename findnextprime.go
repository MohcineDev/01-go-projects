package piscine

func FindNextPrime(nb int) int {
	isPrimeNum := true

	if nb <= 1 {
		isPrimeNum = !isPrimeNum
	} else {
		for i := 2; i <= nb/i; i++ {
			if nb%i == 0 {
				isPrimeNum = false
				break
			}
		}
	}
	if !isPrimeNum {
		return FindNextPrime(nb + 1)
	} else {
		return nb
	}
}

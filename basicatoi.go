package piscine

func BasicAtoi(nbr string) int {
	var result int

	for _, elem := range nbr {

		result *= 10

		if elem >= '0' && elem <= '9' {
			result += int(elem - '0')
		}
	}

	return result
}

package piscine

func CollatzCountdown(start int) int {
	count := 0

	if start <= 0 {
		count = -1
	} else {
		for start != 1 {
			if start%2 == 0 {
				start /= 2
				count++

			} else {
				start = 3*start + 1
				count++
			}
		}
	}

	return count
}

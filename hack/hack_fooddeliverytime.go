package piscine

type food struct {
	preptime int
	name     string
}

func FoodDeliveryTime(order string) int {
	time := 0
	burger := food{
		preptime: 15,
		name:     "burger",
	}
	chips := food{
		preptime: 10,
		name:     "chips",
	}
	nuggets := food{
		preptime: 12,
		name:     "nuggets",
	}

	if order == burger.name {
		time += burger.preptime
	} else if order == chips.name {
		time += chips.preptime
	} else if order == nuggets.name {
		time += nuggets.preptime
	} else {
		time += 404
	}

	return time
}

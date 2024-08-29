package main

func deliSanta(elfDirections string) int {
	numHouses := map[[2]int]bool{}

	var x, y int
	numHouses[[2]int{x, y}] = true

	for _, cha := range elfDirections {
		switch cha {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}

		numHouses[[2]int{x, y}] = true
	}

	return len(numHouses)
}

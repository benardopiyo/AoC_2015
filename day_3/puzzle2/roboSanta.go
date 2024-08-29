package main

// calculates the number of unique houses visited by Santa and Robo-Santa
func roboSanta(elfDirections string) int {
	housesVisited := make(map[[2]int]bool)

	santaX, santaY := 0, 0
	roboX, roboY := 0, 0

	housesVisited[[2]int{santaX, santaY}] = true

	for i, cha := range elfDirections {
		if i%2 == 0 {
			// Santa moves
			switch cha {
			case '^':
				santaY++
			case 'v':
				santaY--
			case '>':
				santaX++
			case '<':
				santaX--
			}
			housesVisited[[2]int{santaX, santaY}] = true
		} else {
			// Robo-Santa moves
			switch cha {
			case '^':
				roboY++
			case 'v':
				roboY--
			case '>':
				roboX++
			case '<':
				roboX--
			}
			housesVisited[[2]int{roboX, roboY}] = true
		}
	}
	return len(housesVisited)
}

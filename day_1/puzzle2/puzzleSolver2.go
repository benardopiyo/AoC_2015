package main

func charPosToBasement(s string) int {
	floor := 0

	for i, cha := range s {
		if cha == '(' {
			floor++
		} else if cha == ')' {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}
	return -1
}

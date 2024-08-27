package main

func NotQuiteLisp(s string) int {
	floor := 0

	for _, cha := range s {
		if cha == '(' {
			floor++
		} else if cha == ')' {
			floor--
		}
	}
	return floor
}

package main

import "sort"

func ribbonWrapper(l, w, h int) int {
	perimeter := 0
	cubic := l * w * h

	sides := []int{l, w, h}
	sort.Slice(sides, func(i, j int) bool {
		return sides[i] < sides[j]
	})

	perimeter = 2 * (sides[0] + sides[1])
	total := perimeter + cubic
	return total
}

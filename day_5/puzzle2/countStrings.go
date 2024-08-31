package main

func countNiceStrings(stringList []string) int {
	count := 0
	for _, str := range stringList {
		if hasSantaRepeatPair(str) && hasSantaSkipOneGap(str) {
			count++
		}
	}
	return count
}
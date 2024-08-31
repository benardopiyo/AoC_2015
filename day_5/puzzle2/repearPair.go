package main

import "strings"

func hasSantaRepeatPair(s string) bool {
	pairIndices := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		lastIdx, found := pairIndices[pair]

		if found && lastIdx < i-1 {
			return true
		}
		pairIndices[pair] = i
	}

	for i := 0; i < len(s)-2; i++ {
		pair := s[i : i+2]
		if j := i + 2; j < len(s)-1 && strings.Contains(s[j:], pair) {
			return true
		}
	}
	return false
}

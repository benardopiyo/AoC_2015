package main

// Checks if the string contains at least one double char
func hasDoubleChar(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if rune(s[i]) == rune(s[i+1]) {
			return true
		}
	}
	return false
}
package main

// Checks if the string contains any of the unwanted pairs
func containsUnwantedPair(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		switch pair {
		case "ab", "cd", "pq", "xy":
			return true
		}
	}
	return false
}
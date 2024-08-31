package main

func isVowel(run rune) bool {
	switch run {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// Checks if the string contains at least three vowels
func hasAtleast3Volels(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if isVowel(rune(s[i])) {
			count++
		}
		if count >= 3 {
			return true
		}
	}
	return false
}

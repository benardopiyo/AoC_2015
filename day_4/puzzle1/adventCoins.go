package main

import (
	"fmt"
	"strings"
)

// Finds the lowest number that,
// when appended to the key, produces a hash starting with five zeroes
func adventCoinSuffix(key string) int {
	number := 1
	for {
		hashInput := fmt.Sprintf("%s%d", key, number)
		hash := md5HashToHex(hashInput)

		if strings.HasPrefix(hash, "00000") {
			return number
		}
		number++
	}
}

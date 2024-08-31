package main

import (
	"crypto/md5"
	"encoding/hex"
)

// Computes the MD5 hash of the input string
// and returns its hexadecimal representation
func md5HashToHex(s string) string {
	hasher := md5.Sum([]byte(s))
	return hex.EncodeToString(hasher[:])
}

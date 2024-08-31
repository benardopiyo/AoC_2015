package main

import (
	"fmt"
)

func main() {
	secretKey := "ckczppom"
	suffix := adventCoinSuffix(secretKey)
	fmt.Printf("%s%d\n", secretKey, suffix)
}

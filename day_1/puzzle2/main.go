package main

import (
	"fmt"
	"log"
)

func main() {
	content, err := ReadFile("file2.txt")
	if err != nil {
		log.Fatalf("Failed to read file %d\n", err)
	}
	fmt.Println(charPosToBasement(content))
}

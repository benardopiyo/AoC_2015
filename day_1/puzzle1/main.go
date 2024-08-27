package main

import (
	"fmt"
	"log"
)

func main() {
	content, err := ReadFile("file1.txt")
	if err != nil {
		log.Fatalf("Failed to read file %d\n", err)
	}

	fmt.Println(NotQuiteLisp(content))
}

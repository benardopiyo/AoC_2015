package main

import (
	"fmt"
	"os"
)

func readFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	 
	return string(data)
}
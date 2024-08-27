package main

import (
	"fmt"
	"os"
)

func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return string(data), nil
}

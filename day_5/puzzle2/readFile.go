package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filePath string) ([]string, error) {
	line, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer line.Close()

	scanner := bufio.NewScanner(line)
	
	var lines []string
	 for scanner.Scan() {
		lines = append(lines, scanner.Text())
	 }

	 if err := scanner.Err(); err != nil {
		return nil, err
	 }

	 return lines, nil
}

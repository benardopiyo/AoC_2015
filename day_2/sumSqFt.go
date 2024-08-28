package main

import (
	"log"
	"strconv"
	"strings"
)

func sumTotal(filePath string) int {
	lines, err := readFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	totalSqFt := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			log.Println("Skipping empty line")
			continue
		}

		dimensions := strings.Split(line, "x")

		if len(dimensions) != 3 {
			log.Fatalf("Invalid dimensions in line: %s", line)
		}

		l, err1 := strconv.Atoi(dimensions[0])
		w, err2 := strconv.Atoi(dimensions[1])
		h, err3 := strconv.Atoi(dimensions[2])

		if err1 != nil || err2 != nil || err3 != nil {
			log.Fatalf("Invalid dimension values in line: %s", line)
		}

		surfaceArea := 2*l*w + 2*w*h + 2*h*l
		minArea := min(l*w, w*h, h*l)
		totalSqFt += surfaceArea + minArea
	}

	return totalSqFt
}

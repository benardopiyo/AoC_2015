package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumTotal(filePath string) int {
	lines, err := readFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	totalSqFt := 0
	for _, line := range lines {
		dimensions := strings.Split(line, "x")

		if len(dimensions) != 3 {
			fmt.Printf("Invalid dimension in line: %s", line)
			continue
		}

		l, err1 := strconv.Atoi(dimensions[0])
		w, err2 := strconv.Atoi(dimensions[1])
		h, err3 := strconv.Atoi(dimensions[2])

		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Printf("Invalid dimension value in line: %s", line)
			continue
		}

		totalSqFt += ribbonWrapper(l, w, h)

	}
	return totalSqFt
}

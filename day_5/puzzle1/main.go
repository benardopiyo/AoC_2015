package main

import (
	"fmt"
	"strings"
)

func main() {
	files := readFile("file5.txt")

	splitFile := strings.Fields(files)

	count := 0
	for _, str := range splitFile {
		if hasAtleast3Volels(str) && hasDoubleChar(str) && !containsUnwantedPair(str) {
			count++
		}
	}
	fmt.Println(count)
}

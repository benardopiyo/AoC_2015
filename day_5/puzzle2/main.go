package main

import (
	"fmt"
)

func main() {
	stringList, err := readFile("file5.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	niceCount := countNiceStrings(stringList)
	fmt.Println(niceCount)
}

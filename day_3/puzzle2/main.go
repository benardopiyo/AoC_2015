package main

import "fmt"

func main() {
	file, err := readFile("file3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(roboSanta(file))
}

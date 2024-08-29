package main

import "fmt"

func main() {
	str, err := readFile("file3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(deliSanta(str))
}

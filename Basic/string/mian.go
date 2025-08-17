package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("String in Go\n")
	str := "apple,banana,orange"

	parts := strings.Split(str, ",")
	fmt.Println("Parts:", parts)
	fmt.Printf("Type of parts: %T\n", parts)

	str2 := "one two one three two one"
	wordCount := strings.Count(str2, "one")
	fmt.Println("Count of 'one':", wordCount)
}


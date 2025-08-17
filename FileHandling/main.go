package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	fmt.Println("File Handling in Go\n")

	//! Input
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close()

	// fmt.Println("File created successfully:", file.Name())
	// content := "Hello, this is a test file2.\n"
	// _, error := io.WriteString(file, content)

	// if error != nil {
	// 	fmt.Println("Error writing to file:", error)
	// 	return
	// }

	//! Output
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, error := io.ReadAll(file)
	if error != nil {
		fmt.Println("Error reading file:", error)
		return
	}
	fmt.Println("File content:\n", string(content))
}
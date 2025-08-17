package main

import (
	"fmt"
)

func main() {
	fmt.Println("Data Conversion in Go\n")

	var intValue int = 42
	var floatValue float64 = 3.14
	var stringValue string = "Hello, Go!"

	var intToFloat float64 = float64(intValue);
	var floatToInt int = int(floatValue)
	var stringToInt int = int(stringValue[0]) // Convert first byte of string to int
	var intToString string = fmt.Sprintf("%d", intValue) // Convert int to string

	fmt.Println("Integer to Float:", intToFloat)
	fmt.Println("Float to Integer:", floatToInt)
	fmt.Println("String to Integer (first byte):", stringToInt)
	fmt.Printf("Integer to String: %T", intToString)
}
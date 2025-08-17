package main

import (
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Started error handling in function in Go");
	result, _ := Divide(10, 0)
	fmt.Println("Result:", result)
}	
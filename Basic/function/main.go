package main

import (
	"fmt"
)

func simpleFunction() {
	fmt.Println("This is a simple function in Go.")
}

func add(a int, b int) (result int) {
	result = a + b
	return;
}

func main() {
	fmt.Println("Learning function in Go");
	simpleFunction();
	result := add(5, 10);
	fmt.Println("Result of addition:", result);
}
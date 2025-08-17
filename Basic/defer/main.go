package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer Eeyword in Go")

	fmt.Println("Starting the main function")
	defer fmt.Println("This will be printed last, after the main function completes")
	defer fmt.Println("Doing some work in the main function")
	fmt.Println("Main function is about to finish")

	// what it does is it defers the execution of the function until the surrounding function returns
	// so defer will execute in LIFO order at the end of the function
}
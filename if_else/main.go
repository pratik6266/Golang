package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning cnditional statements in Go")

	var x int = 20 
	y := 10
	//! can use && || and ! as you know
	if (x > y) {
		fmt.Println("x is greater than y")
	} else if x == y {
		fmt.Println("x is equal to y")	
	} else {
		fmt.Println("x is not greater than y")
	}
}
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

	fmt.Printf("\n-------------------------------------\n\n")

	day := 8;
	switch (day) {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:	
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("Invalid day")
	}
}

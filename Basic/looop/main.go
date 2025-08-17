package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in Go");

	for i:=0; i<10; i++ {
		if i == 0{
			fmt.Println("Zero is neither even nor odd")
		} else if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	//! Infinite loop
	// for{
	// 	fmt.Println("This is an infinite loop, press Ctrl+C to stop it")
	// }
	
	
	// loops in array
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr{
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
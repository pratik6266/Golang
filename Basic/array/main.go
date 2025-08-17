package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array in Go")

	var array[5] int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	fmt.Println("Array elements:", array)

	var names = [5] string{"Pratik", "Raj"};
	fmt.Println("Names in array:", names)
	fmt.Println("Length of names array:", len(names))
	fmt.Println("Capacity of names array:", cap(names))
	//! go does not support multi-dimensional arrays like C, but you can use slices for that
	//! go always intilizes the array with default value of the type

	fmt.Println("\n-------------------------\n")

	var number = []int {1, 2, 3, 4, 5} // slice
	number = append(number, 6, 7, 7)
	fmt.Println("Numbers in slice:", number)
	fmt.Println("Length of number slice:", len(number))
	fmt.Println("Type of number slice:", fmt.Sprintf("%T", number))
	// empty slice can be created like this
	var emptySlice = []int{};
	fmt.Println("Empty slice:", emptySlice)


	var makeSlice = make([]int, 5, 10);
	makeSlice = append(makeSlice, 1, 2, 3, 4,5, 6, 6, 7, 8, 9, 10);
	fmt.Println("Make slice:", makeSlice)
	fmt.Println("Length of make slice:", len(makeSlice))
	fmt.Println("Capacity of make slice:", cap(makeSlice))
}
package main

import (
	"fmt"
)

func changevalue(ptr *int) {
	*ptr = 100 // Change the value at the address pointed to by ptr
}

func main() {
	fmt.Println("Pointer in Go\n")

	var num int = 42

	var ptr *int = &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address stored in ptr:", ptr)
	fmt.Println("Value pointed to by ptr:", *ptr)
	fmt.Println("Address of num:", &num)

	//* just like in C, you can change the value of num using pointer
	changevalue(ptr);
	fmt.Println("Value of num after change:", num)
	fmt.Println("Address of ptr:", &ptr)
	fmt.Println("Value of ptr:", ptr)
	fmt.Println("Value pointed to by ptr after change:", *ptr)
}
package main

import (
	"fmt"
)

type Person struct{
	FirstName string
	LastName string
	Age int
}

func main() {
	fmt.Println("Struct in Go")

	var a Person = Person{
		FirstName: "Pratik",
		LastName: "Raj",
		Age: 21,
	}

	fmt.Println("Person Details:", a)
	fmt.Println("First Name:", a.FirstName)
	fmt.Println("Last Name:", a.LastName)
	fmt.Println("Age:", a.Age)
}
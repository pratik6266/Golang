package main

import (
	"fmt"
	"learning/myUtils"
)

func main() {
	fmt.Println("Hello, World!")
	myUtils.PrintMessage();
	var name string = "Pratik Raj";
	fmt.Println("Hello, " + name + "!")
	var money float64 = 100.50;
	fmt.Println("Money: ", money);
	var decided bool = true;
	fmt.Println("Decision made: ", decided);
	var age int = 25;
	fmt.Println("Age: ", age);
	// Here you can use const as well
	const pi float64 = 3.14; // like this one
	fmt.Println("Value of Pi: ", pi);

	persion := "Pratik Raj" // short variable declaration
	fmt.Println("Person: ", persion)

	//! If your variable first letter is capital then it is exported because its public
	//! If your variable first letter is small then it is not exported because its private
	//todo this is stands for function as well, This is how you export
	//? println in go adds a new line at the end of the string also space between the strings
	//* You can use Printf for formatted output
	fmt.Printf("Hello, %s", "Pratik Raj"); // just like printf in c language
	fmt.Printf("Type of pi varaible is %T\n", pi);
}	
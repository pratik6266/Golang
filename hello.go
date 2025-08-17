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
}	
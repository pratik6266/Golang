package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Printf("Enter your name: ");

	// var name string;
	// fmt.Scan(&name);

	reader := bufio.NewReader(os.Stdin);
	name, _ := reader.ReadString('\n');

	fmt.Printf("Hello, " + name);
}
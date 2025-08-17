package main

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(1 * time.Second) // Simulate some work
	fmt.Println("Hello again after 1 second!")
}

func SayGoodbye() {
	fmt.Println("Goodbye, World!")
}

func main() {
	fmt.Println("Leaning, Go Routines\n");

	go SayHello()
	SayGoodbye()

	time.Sleep(2 * time.Second) // Wait for goroutine to finish
}
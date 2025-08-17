package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Println("Worker", id, "started")
	// Simulate some work
	fmt.Println("Worker", id, "finished")
	defer wg.Done() // Signal that this worker is done
}

func main() {
	fmt.Println("sync in Go")

	var wg sync.WaitGroup
	
	for i:=1; i <= 3;i++ {
		wg.Add(1)
		go worker(i, &wg);
	}

	wg.Wait() 
	fmt.Println("All workers done")
}
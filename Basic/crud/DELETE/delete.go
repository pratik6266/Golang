package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Delete method in Go")

	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
}
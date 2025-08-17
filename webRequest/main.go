package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Fetch in Go, taking response from an API\n")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", res.StatusCode)
		return
	}
	fmt.Printf("Response Type: %T\n", res)
	fmt.Println("Response Status:", res.Status)

	data, errors := io.ReadAll(res.Body)
	if errors != nil {
		fmt.Println("Error reading response body:", errors)
		return
	}
	fmt.Println("Response Body:", string(data))

	fmt.Println("\n---------------------------\n")

	var todo Todo
	err = json.Unmarshal(data, &todo)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Unmarshalled Todo:", todo)
	fmt.Printf("UserID: %d, ID: %d, Title: %s, Completed: %t\n", todo.UserID, todo.ID, todo.Title, todo.Completed)
}
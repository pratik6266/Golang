package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	// "io"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Get method in Go\n")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {		fmt.Println("Error: Status code", res.StatusCode)
		return
	}

	// data, errors := io.ReadAll(res.Body)
	// if errors != nil {
	// 	fmt.Println("Error reading response body:", errors)
	// 	return
	// }

	fmt.Printf("Response Type: %T\n", res)
	fmt.Println("Response Status:", res.Status)
	// fmt.Println("Response: ", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {	
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Todo Item:")
	fmt.Printf("UserID: %d, ID: %d, Title: %s, Completed: %t\n", todo.UserID, todo.ID, todo.Title, todo.Completed)
	fmt.Println("Data:", todo)
}
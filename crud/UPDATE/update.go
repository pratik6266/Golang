package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Update method in Go\n")

	var todo = Todo{
		UserID:    27,
		ID:        27,
		Title:     "Updated Todo",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("JSON Output:", string(jsonData))

	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/27", strings.NewReader(string(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making PUT request:", err)
		return
	}

	var resData Todo
	err = json.NewDecoder(res.Body).Decode(&resData)
	if err != nil {
		fmt.Println("Error decoding response JSON:", err)
		return
	}

	fmt.Println("Response from server:", resData)
	fmt.Println("UserID:", resData.UserID)
	fmt.Println("ID:", resData.ID)
	fmt.Println("Title:", resData.Title)
	fmt.Println("Completed:", resData.Completed)
	fmt.Printf("Type of response: %T\n", resData)
}
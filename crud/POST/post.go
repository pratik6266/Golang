package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
)

type Todos struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Post method in Go\n")

	var todo = Todos{
		UserID:    2,
		ID:        2,	
		Title:     "Sample Todo",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println("JSON Output:", string(jsonData))

	res, errors := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", strings.NewReader(string(jsonData)))

	if errors != nil {
		fmt.Println("Error making POST request:", errors)
		return
	}
	defer res.Body.Close()

	var resData Todos
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
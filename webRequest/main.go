package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

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

	data, errors := ioutil.ReadAll(res.Body)
	if errors != nil {
		fmt.Println("Error reading response body:", errors)
		return
	}
	fmt.Println("Response Body:", string(data))
}
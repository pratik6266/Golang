package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL in Go\n")

	URL := "https://api.example.com/data?key=1"
	
	parsedUrl, err := url.Parse(URL)
	if err != nil {		
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Printf("Type of parsedUrl: %T\n", parsedUrl)
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Raw Query:", parsedUrl.RawQuery)
}
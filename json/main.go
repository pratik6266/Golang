package main

import (
	"fmt"
	"encoding/json"
)

type Person struct{
	Name string `json:"name"`
	Age  int `json:"age"`
}

func main() {
	fmt.Println("Learning URL in Go\n")

	me := Person{
		Name: "Pratik",
		Age:  25,
	}

	fmt.Printf("Type of me: %T\n", me)
	fmt.Printf("Name: %s, Age: %d\n", me.Name, me.Age)
	fmt.Println("Data: ", me)

	// convert person to JSON(Marshalling)
	data, err := json.Marshal(me)
	if err != nil {		
		fmt.Println("Error marshalling data:", err)
		return
	}

	fmt.Println("JSON data:", string(data))

	// convert JSON back to person (Unmarshalling)
	var person Person
	err = json.Unmarshal(data, &person)
	if err != nil {	
		fmt.Println("Error unmarshalling data:", err)
		return
	}

	fmt.Println("Unmarshalled Person:", person)
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	fmt.Println("Data:", person)
}
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go are key-value pairs.")

	studentGrades := make(map[string]int)
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 92

	fmt.Println("Student Grades:")
	for key, value := range studentGrades {
		fmt.Printf("%s, Grade: %d\n", key, value)
	}
	delete(studentGrades, "Alice")
	//default value of any is 0 in map

	//check if a key exists
	grade, exits := studentGrades["Bob"]
	fmt.Println("Grade of Bob:", grade, "Exists:", exits)
}
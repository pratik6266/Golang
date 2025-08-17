package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in Go\n")

	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Println("Year:", currentTime.Year())
	fmt.Println("Month:", currentTime.Month())
	fmt.Println("Day:", currentTime.Day())
	fmt.Println("Hour:", currentTime.Hour())
	fmt.Println("Minute:", currentTime.Minute())
	fmt.Println("Second:", currentTime.Second())
	fmt.Println("Location:", currentTime.Location())
	fmt.Println("Unix Timestamp:", currentTime.Unix())
	fmt.Println("Formatted Time:", currentTime.Format("2006 01 02"))
	fmt.Println("Formatted Time:", currentTime.Format("15:04:05"))
	fmt.Println("Type: ", fmt.Sprintf("%T", currentTime))
	fmt.Println("Day: ", currentTime.Format("Monday"))
}
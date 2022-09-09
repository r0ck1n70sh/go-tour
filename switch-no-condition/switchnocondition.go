package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Greetings...")

	hour := time.Now().Hour()
	// fmt.Printf("hour -> %d\n", hour)

	switch {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Night!")
	}
}

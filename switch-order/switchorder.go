package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today:
		fmt.Println("today")

	case today + 1:
		fmt.Println("tomorrow")

	case today + 2:
		fmt.Println("day after tomorrow")

	default:
		fmt.Println("after a long time...")
	}
}

package main

import (
	"fmt"
	"math"
)

func sqrt(num float64) string {
	if num < 0 {
		return sqrt(-num) + "i"
	}

	return fmt.Sprintf("%f", math.Sqrt(num))
}

func main() {
	num1 := 2.0
	num2 := -2.0

	fmt.Println("num1", sqrt(num1))
	fmt.Println("num2", sqrt(num2))
}

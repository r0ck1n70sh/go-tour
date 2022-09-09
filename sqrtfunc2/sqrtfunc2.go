package main

import (
	"fmt"
	"math"
)

func sqrt(num float64) string {
	if num < 0 {
		return sqrt(-num) + "i"
	}

	return fmt.Sprintf("%.2f", math.Sqrt(num))
}

func main() {
	num := -4.0
	fmt.Printf("sqrt(%.2f): %s\n", num, sqrt(num))
}

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64, x, y float64) float64 {
	return fn(x, y)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Printf("compute(hypot, 5, 12): %f\n", compute(hypot, 5, 12))
	fmt.Printf("compute(hypot, 3, 4): %f\n", compute(hypot, 3, 4))
}

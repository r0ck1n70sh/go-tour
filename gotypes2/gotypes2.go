package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y = 1, 2
	var sqroot float64 = -math.Sqrt(float64(x*x + y*y))
	var v = uint(sqroot)

	fmt.Println(x, y, sqroot, v)
}

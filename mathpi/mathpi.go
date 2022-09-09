package main

import (
	"fmt"
	"math"
)

func fetchPi() float32 {
	pi := float32(math.Pi)
	return pi
}

func main() {
	pi := fetchPi()
	fmt.Println(pi)
}

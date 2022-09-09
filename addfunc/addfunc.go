package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func main() {
	a, b := 12, 8
	sum := add(a, b)

	fmt.Println("sum:", sum)

	sum2 := add2(a, b)
	fmt.Println("sum2:", sum2)
}

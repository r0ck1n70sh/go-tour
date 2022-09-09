package main

import "fmt"

func main() {
	sum := 1
	n := 0

	for sum < 1000 {
		sum += sum
		n++
	}

	fmt.Println("sum:", sum, "n:", n)
}

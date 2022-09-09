package main

import (
	"fmt"
)

func concat(x, y string) string {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := "abc", "xyz"

	orig := concat(a, b)
	rev := concat(swap(a, b))

	fmt.Println("orig:", orig)
	fmt.Println("rev:", rev)
}

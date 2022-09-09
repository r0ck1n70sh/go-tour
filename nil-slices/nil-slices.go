package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("s: %v, len(s):%d, cap(s):%d\n", s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}

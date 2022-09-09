package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("b[:2]", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(x string, s []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", x, len(s), cap(s), s)
}

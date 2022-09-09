package main

import "fmt"

func main() {
	//create array, then builds a slice for it

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("q:", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("r:", r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("s:", s)
}

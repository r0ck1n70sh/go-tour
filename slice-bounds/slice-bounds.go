package main

import "fmt"

func main() {
	s := []int{2, 3, 7, 11, 13}
	fmt.Println("s", s)

	// s[default=0: defualt=length]
	a := s[1:4]
	fmt.Println("s[1:4]", a)

	b := s[:2]
	fmt.Println("s[:2]", b)

	c := s[1:]
	fmt.Println("s[1:]", c)

	d := s[:]
	fmt.Println("s[:]", d)
}

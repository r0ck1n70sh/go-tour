package main

import "fmt"

func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println("names:", names)

	a := names[0:2]
	fmt.Println("a", a)

	b := names[1:3]
	fmt.Println("b", b)

	b[0] = "XXX"
	fmt.Println("setting, b[0]=\"XXX\"")

	fmt.Println("a", a)
	fmt.Println("b", b)

	fmt.Println("names", names)
}

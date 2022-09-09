package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	fmt.Printf("pow %v\n", pow)

	for _, val := range pow {
		fmt.Printf("%d ", val)
	}

	fmt.Printf("\n")
}

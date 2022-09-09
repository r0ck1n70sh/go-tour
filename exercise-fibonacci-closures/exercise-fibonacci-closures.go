package main

import "fmt"

func fibonacci() func() int {
	fib1 := -1
	fib2 := 1
	return func() int {
		temp := fib2
		fib2 = fib1 + fib2
		fib1 = temp
		return fib2
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Printf("fibonacci[\"%d\"]: %d\n", i, f())
	}
}

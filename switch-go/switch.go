package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go running on...")
	os := runtime.GOOS
	// os = "Free BSD"

	switch os {
	case "darwin":
		fmt.Println("OS X")

	case "linux":
		fmt.Println("Linux")

	default:
		fmt.Printf("%s. \n", os)
	}
}

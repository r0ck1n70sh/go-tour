package main

import (
	"fmt"
	"time"
)

func main() {
	milli := timestamp()

	fmt.Println("Generating a random number:", milli)

	timeMilliseconds()
}

func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func timeMilliseconds() {
	fmt.Println("typeOf time.Millisecond", int32(time.Millisecond))
}

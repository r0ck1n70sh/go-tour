package main

import "fmt"

type Vertex struct {
	X, Y float64
}

var m = map[string]Vertex{
	"bell_labs": Vertex{45.7482, -67.3438},
	"google":    Vertex{-89.3243, 23.5865},
}

func main() {
	fmt.Println("m", m)
}

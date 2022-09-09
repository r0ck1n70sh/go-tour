package main

import "fmt"

type Vertex struct {
	X, Y float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["home"] = Vertex{40.45387, -56.34398}

	fmt.Println("m[\"home\"]:", m["home"])
}

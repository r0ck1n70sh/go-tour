package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) dist(p1 Point) float64 {
	diffX, diffY := p.x-p1.x, p.y-p1.y

	return math.Sqrt(diffX*diffX + diffY*diffY)
}

func main() {
	a, b := Point{1.4, -7.7}, Point{-1.6, 11.3}
	fmt.Printf("dist(a,b) %.2f\n", a.dist(b))
}

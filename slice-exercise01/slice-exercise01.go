package main

import "golang.org/x/tour/pic"

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for i := range pic {
		pic[i] = make([]uint8, dx)
	}

	return pic
}

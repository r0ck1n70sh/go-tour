package main

import "fmt"

func main() {
	m := make(map[string]int)
	key := "answer"

	m[key] = 42
	printM(m, key)

	m[key] = 48
	printM(m, key)

	delete(m, key)
	printM(m, key)

	val, isPresent := m[key]
	fmt.Printf("val:%d, isPresent:%t\n", val, isPresent)

}

func printM(m map[string]int, key string) {
	fmt.Printf("m[\"%s\"]: %d\n", key, m[key])
}

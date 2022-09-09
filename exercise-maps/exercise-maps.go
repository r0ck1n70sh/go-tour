package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCounts(s string) map[string]int {
	str := strings.Split(s, " ")
	m := make(map[string]int)

	for i := 0; i < len(str); i++ {
		ch := string(str[i])
		m[ch] = m[ch] + 1
	}

	return m
}

func main() {
	wc.Test(WordCounts)
}

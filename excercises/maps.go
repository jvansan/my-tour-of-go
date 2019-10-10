package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount takes string and counts unique words
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		v, ok := m[word]
		if ok != true {
			m[word] = 1
		} else {
			m[word] = v + 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

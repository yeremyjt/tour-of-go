package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	words := strings.Split(s, " ")

	for _, word := range words {
		elem, ok := m[word]
		if ok == true {
			m[word] = elem + 1
		} else {
			m[word] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}

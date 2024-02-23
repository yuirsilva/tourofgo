package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	word_map := make(map[string]int, 10)
	res := strings.Fields(s)
	for _, el := range res {
		word_map[el]++
	}

	return word_map
}

func main() {
	wc.Test(WordCount)
}

// https://go.dev/tour/moretypes/23

package main

import "fmt"

func main() {
	// another way to say "any"
	// every method implements at least zero methods
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// https://go.dev/tour/methods/14

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	value := rand.Intn(10)
	if value == 1 {
		fmt.Printf("Now you have %v problem.\n", value)
	} else {
		fmt.Printf("Now you have %v problems.\n", value)
	}
}

// https://go.dev/tour/basics/2

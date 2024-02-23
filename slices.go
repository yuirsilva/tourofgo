package main

import "fmt"

func main() {
	// slices == array
	// slices does not store any data
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

// https://go.dev/tour/moretypes/7

package main

import "fmt"

func main() {
	sum := 1
	// NO WHILE
	// WHILE == FOR (no init or post)
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// https://go.dev/tour/flowcontrol/3

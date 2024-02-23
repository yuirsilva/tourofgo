package main

import "fmt"

var pow = []int{1, 1, 2, 3, 5, 8, 13, 21}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// https://go.dev/tour/moretypes/16

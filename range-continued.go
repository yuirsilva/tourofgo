package main

import (
	"fmt"
	"math"
)

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = int(math.Pow(2, float64(i))) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// https://go.dev/tour/moretypes/17

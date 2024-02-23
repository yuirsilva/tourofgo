package main

import (
	"fmt"
	"math"
)

// simple type
type MyFloat float64

// assigning methods to types
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// https://go.dev/tour/methods/3

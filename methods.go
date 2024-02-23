package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// methods are functions :D
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// https://go.dev/tour/methods/1

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// methods are functions :D
// value receiver
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

// https://go.dev/tour/methods/2

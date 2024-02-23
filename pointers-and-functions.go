package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// pointer in argument
	// pointer to v Vertex
	// methods are functions but functions
	// behave differently
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

// https://go.dev/tour/methods/5

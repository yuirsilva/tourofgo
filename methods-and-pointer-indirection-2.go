package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	// you can pass a pointer to a value receiver (methods)
	fmt.Println(p.Abs())
	// you have to pass a value to a value argument (functions)
	// you can't pass a pointer to a value argument (functions)
	fmt.Println(AbsFunc(*p))
}

// https://go.dev/tour/methods/7

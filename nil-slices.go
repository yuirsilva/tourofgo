package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil slice! No underlying array")
		fmt.Println("capacity and length = 0")
	}
}

// https://go.dev/tour/moretypes/12

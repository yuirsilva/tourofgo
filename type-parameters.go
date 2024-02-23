package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 15, 20}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "bars"}
	fmt.Println(Index(ss, "lmao"))
}

// https://go.dev/tour/generics/1

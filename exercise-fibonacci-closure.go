package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// starts from 0 and 1
	f0, f1 := 0, 1
	return func() int {
		f2 := f0
		f0, f1 = f1, f2+f1

		return f2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// https://go.dev/tour/moretypes/26

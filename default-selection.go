package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	select {
	case ch <- 666:
		fmt.Println("too bad!")
		// this case would block
		// (no other routine receives the value)
	case i, ok := <-ch:
		fmt.Println("too bad!", i, ok)
		// this case would block
		// (no other routine sends a value)
	default:
		fmt.Println("good! no locks!")
		// this case would not block
		// (no other case is ever ready, so this wouldn't block)
		// (nor is this case receiving/sending)
	}
}

// https://go.dev/tour/concurrency/6

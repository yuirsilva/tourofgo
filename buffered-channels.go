package main

import "fmt"

func main() {
	// sends get blocked when the buffer is full
	// receives get blocked when the buffer is empty
	ch := make(chan int, 1)
	ch <- 4
	fmt.Println("Hello")
	fmt.Println(<-ch)
}

// https://go.dev/tour/concurrency/3

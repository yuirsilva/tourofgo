package main

import "fmt"

// func never(ch chan int) {
// 	x := 0
// 	fmt.Println("sent1")
// 	select {
// 	case ch <- x:
// 		fmt.Println("sent2")
// 	}
// 	fmt.Println("sent3")
// }

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// https://go.dev/tour/concurrency/5

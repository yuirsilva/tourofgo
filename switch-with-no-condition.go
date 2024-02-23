package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// if then
	switch {
	case t.Hour() < 12:
		fmt.Println("Good manhã.")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// https://go.dev/tour/flowcontrol/11

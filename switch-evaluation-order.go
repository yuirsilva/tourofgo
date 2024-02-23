package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today :D")
	case today + 1:
		fmt.Println("Tomorrow :D")
	case today + 2:
		fmt.Println("In two days f")
	default:
		fmt.Println("Too far away pepesad")
	}
}

// https://go.dev/tour/flowcontrol/10

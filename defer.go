package main

import "fmt"

func main() {
	defer fmt.Println("end")

	fmt.Println("Hello")
}

// https://go.dev/tour/flowcontrol/12

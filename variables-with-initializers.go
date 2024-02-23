package main

import "fmt"

var i, j int = 1, 2

func main() {
	var htmx, python, java = true, !false, "no!"
	fmt.Println(i, j, htmx, python, java)
}

// https://go.dev/tour/basics/9

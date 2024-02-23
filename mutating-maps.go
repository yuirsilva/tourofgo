package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["BestEditor"] = "Vim"
	fmt.Println("The value:", m["BestEditor"])

	m["BestLanguage"] = "Go"
	fmt.Println("The value:", m["BestLanguage"])

	delete(m, "BestEditor")
	fmt.Println("The value:", m["BestEditor"])

	v, ok := m["BestEditor"]
	fmt.Println("The value:", v, "Present?", ok)
}

// https://go.dev/tour/moretypes/22

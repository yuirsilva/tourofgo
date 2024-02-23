package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Sihat Labs": {
		666.666, -74.39967,
	},
	"BNYX": {
		37.42202, -122.30,
	},
}

func main() {
	fmt.Println(m)
}

// https://go.dev/tour/moretypes/20

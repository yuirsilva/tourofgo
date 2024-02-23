package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	for tmp := 1.0; ; {
		state := tmp
		tmp -= (tmp*tmp - x) / (2 * tmp)

		if float32(state) == float32(tmp) {
			return tmp
		}
	}
}

func main() {
	fmt.Println(Sqrt(16))
}

// https://go.dev/tour/flowcontrol/8

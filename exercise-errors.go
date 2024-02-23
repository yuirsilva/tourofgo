package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Can not Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	for tmp := 1.0; ; {
		state := tmp
		tmp -= (tmp*tmp - x) / (2 * tmp)

		if float32(state) == float32(tmp) {
			return tmp, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(16))
	fmt.Println(Sqrt(-16))
}

// https://go.dev/tour/methods/20

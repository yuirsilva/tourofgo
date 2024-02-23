package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	file := strings.NewReader("A file is a reader!")

	b := make([]byte, 5)

	for {
		n, err := file.Read(b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// https://go.dev/tour/methods/21

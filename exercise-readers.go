package main

import (
	"io"
)

type MyReader struct {
	// s rune
}

func (r MyReader) Read(p []byte) (n int, err error) {
	if n > len(p) {
		return 0, io.EOF
	}
	for i := range p {
		// p[i] = byte(r.s)
		p[i] = byte('A')
		n += 1
	}

	return
}

func main() {
	// buffer := make([]byte, 12)

	// my := MyReader{s: 'A'}
	// n, err := my.Read(buffer)
	// fmt.Println(n, err)
}

// https://go.dev/tour/methods/22

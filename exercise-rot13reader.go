package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i := range b[:n] {
		if b[i] >= 65 && b[i] <= 90 {
			b[i] = ((b[i]%65 + 1 + 13) % 26) + 64
		} else if b[i] >= 97 && b[i] <= 122 {
			b[i] = ((b[i]%97 + 1 + 13) % 26) + 96
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rot := rot13Reader{s}
	io.Copy(os.Stdout, &rot)
}

// https://go.dev/tour/methods/23

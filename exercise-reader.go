package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	if x >= 'A' && x <= 'M' {
		return x + 13
	}

	if x >= 'N' && x <= 'Z' {
		return x - 13
	}

	if x >= 'a' && x <= 'm' {
		return x + 13
	}

	if x >= 'n' && x <= 'z' {
		return x - 13
	}

	return x

}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := range b {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rt *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rt.r.Read(b)
	for i := 0; i < len(b); i++ {
		a := b[i]
		if (a >= 'a' && a <= 'm') || (a >= 'A' && a <= 'M') {
			b[i] += 13
		} else if (a >= 'n' && a <= 'z') || (a >= 'N' && a <= 'Z') {
			b[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

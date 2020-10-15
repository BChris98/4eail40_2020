package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (r spaceEraser) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	j := 0
	for i := 0; i < n; i++ {
		if p[i] != 32 {
			p[j] = p[i]
			j++
		}
	}
	return j, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

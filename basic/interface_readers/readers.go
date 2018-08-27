// examples of io.Reader
//func (T) Read(b []byte) (n int, err error)

package main

import (
	. "fmt"
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	
	var n int
	var err error

	n, err = r.r.Read(b)

	for index, value := range b[:n] {
		if value <= 78 {
			b[index] = value + 13
		}
		if value <=90 {
			b[index] = value - 14
		}
		if value <= 110 {
			b[index] = value + 13
		}
		if value <= 122 {
			b[index] = value - 14
		}
	}
	
	return n, err
}

func main() {
	r := strings.NewReader("hello, world")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)	
		Printf("n: %v, err: %v, b: %v\n", n, err, b)
		Printf("b[:n] is %v\n", b[:n])
		if err == io.EOF {
			break
		}
	}	

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r2 := rot13Reader{s}
	io.Copy(os.Stdout, &r2)
}





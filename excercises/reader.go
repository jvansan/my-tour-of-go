package main

import "golang.org/x/tour/reader"

// MyReader tutorial struct
type MyReader struct{}

// Read implements Reader specified at:
// https://golang.org/pkg/io/#Reader
func (r MyReader) Read(s []byte) (i int, err error) {
	size := len(s)
	for i := 0; i < size; i++ {
		s[i] = 'A'
	}
	return size, nil
}

func main() {
	reader.Validate(MyReader{})
}

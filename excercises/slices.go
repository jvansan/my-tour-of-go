package main

import "golang.org/x/tour/pic"

// Pic takes input from tour and returns 2D array
func Pic(dx, dy int) [][]uint8 {
	// Allocate two-dimensioanl array.
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	for idy := range a {
		for idx := range a[idy] {
			a[idy][idx] = uint8(idx * idy)
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}

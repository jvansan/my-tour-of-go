package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// While loop implementation of Newton's method
	// for calculation Sqrt
	z := 1.0
	newZ := z - (z*z-x)/(2*z)

	for math.Abs(z-newZ) > 1E-10 {
		z = newZ
		newZ = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(math.Sqrt(4))
}

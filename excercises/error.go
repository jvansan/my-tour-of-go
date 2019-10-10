package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt is a type for catching negative numbers in Sqrt
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt function from basic but catches negative number and returns error
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	newZ := z - (z*z-x)/(2*z)

	for math.Abs(z-newZ) > 1e-9 {
		z = newZ
		newZ = z - (z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

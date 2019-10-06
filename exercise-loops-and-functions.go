package main

import (
	"fmt"
	"math"
)

func is_converged(x, z float64) bool {
	y := math.Pow(z, 2) - x
	if y < 0 {
		y = -y
	}
	if y < 0.0000000001 {
		return true
	} else {
		return false
	}
}

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 10; i++ {
		if is_converged(x, z) {
			fmt.Println(i)
			fmt.Println(math.Sqrt(x))
			return z
		}
		z -= (z*z - x) / (2 * z)
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(1234))
}

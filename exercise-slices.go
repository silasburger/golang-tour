package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		y[i] = make([]uint8, dx)
	}
	for ai := range y {
		for bi, bv := range y[ai] {
			y[ai][bi] = bv + 2*uint8(bi)
		}
	}
	return y
}

func main() {
	pic.Show(Pic)
}

package main

import (
	"fmt"
	"math"
)

func MySqrt(a float64) (x float64, y string) {
	if a < 0 {
		y = "error"
	}
	x = math.Sqrt(a)
	return
}

func main() {
	a := MySqrt(-1)
	fmt.Println(a)
}

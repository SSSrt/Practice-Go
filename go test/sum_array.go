package main

import (
	"fmt"
)

func main() {
	a := []float32{1.2, 2.3, 4.5, 3.8}
	fmt.Println(Sum(a[:]))
}

func Sum(a []float32) float32 {
	var x float32
	for i := 0; i < len(a); i++ {
		x = x + a[i]
	}
	return x
}

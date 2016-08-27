package main

import (
	"fmt"
)

func sum1(a, b int) (x, y, z int) {
	x = a + b
	y = a * b
	z = a - b
	return x, y, z
}

func sum2(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func main() {
	a, b, c := sum1(5, 6)
	fmt.Println(a, b, c)
	x, y, z := sum2(7, 8)
	fmt.Println(x, y, z)
}

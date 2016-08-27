package main

import "fmt"

func main() {
	a, b := fibonacci(10)
	fmt.Println(a, b)
}

func fibonacci(a int) (x, y int) {
	x = 0
	y = 1
	var z int
	for i := 0; i <= a; i++ {
		z = y
		y = fibonacci1(z)
		x++
	}
	x -= 1
	return x, y
}

func fibonacci1() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

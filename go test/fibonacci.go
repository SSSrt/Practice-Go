package main

import (
	"fmt"
)

func main() {
	a, b := fibonacci(10)
	fmt.Println(a, b)
}

func fibonacci(a int) (x, y int) {
	x = 0
	for i := 0; i <= a; i++ {
		y = fibonacci1(i)
		x += 1
	}
	x -= 1
	return x, y
}

func fibonacci1(a int) (res int) {
	if a <= 1 {
		res = 1
	} else {
		res = fibonacci1(a-1) + fibonacci1(a-2)
	}
	return
}

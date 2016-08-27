package main

import (
	"fmt"
)

func minSlice(a []int) int {
	x := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			x = a[i]
		}
	}
	return x
}

func maxSlice(a []int) int {
	x := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[i-1] {
			x = a[i]
		}
	}
	return x
}

func main() {
	var arr = [5]int{0, 1, 2, 3, 4}
	fmt.Println(minSlice(arr[:]))
	fmt.Println(maxSlice(arr[:]))
}

package main

import (
	"fmt"
)

func main() {
	factor := 10
	a := make([]int, 10)
	newSlice := make([]int, len(a)*factor)
	copy(newSlice, a)
	a = newSlice
	fmt.Println(len(a))

	s := []int{1, 2, 3, 4, 5}
	s1 := []int{7, 8, 9}
	s = append(s[:1], append(s1, s[1:]...)...)
	fmt.Println(s)
	x := []string{"start", "an", "apple", "end"}
	x = append(x[:1], x[3:]...)
	fmt.Println(x)
}

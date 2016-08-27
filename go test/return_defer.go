package main

import "fmt"

func c() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func main() {
	fmt.Println(c())
}

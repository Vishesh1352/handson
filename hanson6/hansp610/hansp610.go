package main

import (
	"fmt"
)

func main() {
	a := foo(10)
	fmt.Println(a)
}

func foo(n int) int {
	var x int
	x = n
	return x
}

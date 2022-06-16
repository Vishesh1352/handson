package main

import (
	"fmt"
)

func main() {
	ni := []int{10, 20, 30, 40, 50, 60}
	a := foo(ni...)
	b := bar(ni)
	fmt.Println("sum in foo", a)
	fmt.Println("sum in bar", b)
}
func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

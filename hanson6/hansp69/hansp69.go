package main

import (
	"fmt"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	a1 := sum(s1...)
	a2 := bar(sum, s1...)
	fmt.Println(a1)
	fmt.Println(a2)

}
func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return (total)
}
func bar(f func(xi ...int) int, li ...int) int {
	sum := 0
	for _, v := range li {
		sum += v
	}
	return sum
}

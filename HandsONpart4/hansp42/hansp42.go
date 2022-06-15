package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)
}

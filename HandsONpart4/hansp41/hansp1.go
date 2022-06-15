package main

import "fmt"

func main() {
	a := [5]int{10, 20, 30, 40}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)
}

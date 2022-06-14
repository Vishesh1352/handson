package main

import "fmt"

type rum int

var x rum

func main() {
	fmt.Print(x)
	fmt.Printf("\n%T\n", x)
	x = 42
	fmt.Print(x)
}

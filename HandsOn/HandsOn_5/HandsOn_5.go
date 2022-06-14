package main

import "fmt"

type rum int

var x rum
var y int

func main() {
	fmt.Print(x)
	fmt.Printf("\n%T\n", x)
	x = 42
	fmt.Print(x)
	y = int(x)
	y = 43
	fmt.Printf("\n")
	fmt.Print(y)
}

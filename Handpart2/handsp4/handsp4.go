package main

import "fmt"

var x int = 11

func main() {
	fmt.Printf("%d\t\t%b\t\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", y, y, y)
}

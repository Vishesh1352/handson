package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "james bond"
	z = true
	s := fmt.Sprint(x, y, z)
	fmt.Print(s)
}

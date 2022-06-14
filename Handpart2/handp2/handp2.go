package main

import "fmt"

var a int = 10
var b int = 555

func main() {
	fmt.Println(a)
	fmt.Printf("\n")
	fmt.Println(b)
	fmt.Printf("\n")
	c := (a == b)
	d := (a <= b)
	e := (a >= b)
	f := (a != b)
	g := (a < b)
	h := (a > b)
	fmt.Println(c)
	fmt.Printf("\n")
	fmt.Println(d)
	fmt.Printf("\n")
	fmt.Println(e)
	fmt.Printf("\n")
	fmt.Println(f)
	fmt.Printf("\n")
	fmt.Println(g)
	fmt.Printf("\n")
	fmt.Println(h)
	fmt.Printf("\n")
}

package main

import "fmt"

const (
	a     = 10 // untype constant
	b     = 0.25
	c     = "hands on practice"
	e int = 55 //type constnt
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", e)
}

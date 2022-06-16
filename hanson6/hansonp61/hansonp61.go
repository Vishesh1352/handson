package main

import (
	"fmt"
)

func main() {
	a := foo()
	b, c := bar()

	fmt.Println("from foo", a)
	fmt.Println("from bar", b)
	fmt.Println("from bar ", c)
}

func foo() int {
	return 54
}
func bar() (int, string) {
	return 54, "hello"
}

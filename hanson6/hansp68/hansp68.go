package main

import (
	"fmt"
)

func main() {
	g := bar()
	fmt.Println(g())
}

func bar() func() int {
	return func() int {
		return 56
	}
}

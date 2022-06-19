package main

import (
	"fmt"
)

type customErr struct {
	info string
	err  error
}

func (ce customErr) Error() string {
	return fmt.Sprintf("somre error happen %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more coffe",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran ", e, "\n")
}

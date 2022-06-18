package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int, 1)
	go func() {
		c1 <- 44
	}()
	fmt.Println(<-c1)
	c2 <- 55
	fmt.Println(<-c2)
}

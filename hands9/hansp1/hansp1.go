package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("printing in main")
	wg.Wait()
}
func foo() {
	fmt.Println("go 1")
	wg.Done()
}
func bar() {
	fmt.Println("go 3")
	wg.Done()
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	counter := 0
	const cg int = 100
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("no go run", runtime.NumGoroutine())
	wg.Add(cg)
	for i := 0; i < cg; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			fmt.Println("no go run", runtime.NumGoroutine())
			mu.Unlock()
			wg.Done()
		}()
	}

	fmt.Println(counter)
	fmt.Println("no go run", runtime.NumGoroutine())
	wg.Wait()
}

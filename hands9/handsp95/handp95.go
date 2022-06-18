package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var counter int64
	const cg int = 100
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("no go run", runtime.NumGoroutine())

	wg.Add(cg)
	for i := 0; i < cg; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	fmt.Println(counter)
	fmt.Println("no go run", runtime.NumGoroutine())
	wg.Wait()
}

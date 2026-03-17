package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count uint64
	var wg sync.WaitGroup

	for i := 0; i < 100_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&count, 1)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

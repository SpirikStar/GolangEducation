package main

import (
	"fmt"
	"sync"
)

func worker(id uint, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Printf("Worker %d; Value %d\n", id, n)
	}
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(2)
	go worker(1, ch, &wg)
	go worker(2, ch, &wg)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
	fmt.Println(ch)
}

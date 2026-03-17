package main

import (
	"fmt"
	"sync"
)

func cacl(mutex *sync.Mutex, number *uint, count int, wg *sync.WaitGroup) {
	mutex.Lock()
	*number += 1
	mutex.Unlock()
	
	fmt.Println("func:", count)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var number uint

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go cacl(&mutex, &number, i, &wg)
	}

	wg.Wait()
	fmt.Println(number)
}

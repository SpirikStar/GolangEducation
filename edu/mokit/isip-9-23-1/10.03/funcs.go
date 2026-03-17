package main

import (
	"fmt"
	"sync"
)

func actionFunc(count *int) {
	*count++
}

func actionWorker(id uint16, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker id:", id)
}

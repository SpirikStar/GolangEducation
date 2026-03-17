package main

import (
	"fmt"
	"sync"
)

func worker(count *uint16, wg *sync.WaitGroup) {
	*count++
	fmt.Println(*count)
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	var count uint16

	index := 0
	for index <= 10 {
		index++
		wg.Add(1)
		go worker(&count, &wg)
	}
	wg.Wait()
	fmt.Println(count)
}

// TODO: Создать горутину через WaitGroup
// TODO: которая вносит изменения в []int{}
// TODO: Функция должна принимать []int{}
// TODO: и добавлять случайно-уникальное число
// TODO: Реализовать управление потоками


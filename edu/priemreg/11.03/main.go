package main

import (
	"fmt"
	"time"
)

func worker(id *uint) {
	*id++
}
func main() {
	var index uint
	var count uint
	for index < 1000 {
		index++
		go worker(&count)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(count)
}

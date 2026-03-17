package main

import (
	"fmt"
	"time"
)

func workerOne(count *uint16) {
	*count++
}

func main() {
	var count uint16
	index := 0
	for index < 1000 {
		index++
		go workerOne(&count)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(count)
}

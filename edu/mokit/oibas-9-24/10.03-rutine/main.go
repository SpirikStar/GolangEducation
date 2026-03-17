package main

import (
	"fmt"
	"time"
)

func funcStart() {
	fmt.Println("Hello Go!")
}

// https://metanit.com/go/tutorial/7.1.php
// https://metanit.com/go/tutorial/7.2.php
// https://metanit.com/go/tutorial/7.3.php


func worker(id int) {
	fmt.Println("Операция (start):", id)
	time.Sleep(1 * time.Second)
	fmt.Println("Операция (stop):", id)
}
func main() {
	count := 0
	for count < 5 {
		count++
		go worker(count)
	}
	go funcStart()
	time.Sleep(2 * time.Second)
}

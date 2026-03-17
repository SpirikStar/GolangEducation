package main

import (
	"fmt"
	"time"
)

func action(val int) {
	fmt.Println("Func action:", val)
}

func main() {
	go action(1)
	go action(2)
	go action(3)
	time.Sleep(1 * time.Second)
}


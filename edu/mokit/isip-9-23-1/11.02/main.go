package main

import (
	"fmt"
	"time"
)

func main() {
	funcOne()
	funcTwo("is", "words")

	count := 0
	funcThree(&count)
	fmt.Println(count)
	
	dTime := time.Now()
	dTimeUnix := time.Now().UnixNano()
	fmt.Println(dTime, dTimeUnix)
}

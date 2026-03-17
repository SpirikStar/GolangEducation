// go mod init lesson
// go run .
package main

import (
	// "fmt"
	// "time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// count := 0
	var index uint16
	for index < 10 {
		index++
		wg.Add(1)
		go actionWorker(index, &wg)
		// 	go actionFunc(&count)
	}
	wg.Wait()
	// time.Sleep(1 * time.Second)
	// fmt.Println(count)

}

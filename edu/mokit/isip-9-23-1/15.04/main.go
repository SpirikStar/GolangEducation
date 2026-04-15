// https://avatars.githubusercontent.com/u/
// TODO: go mod init github
// TODO: go run .
package main

import (
	"log"
	"os"
	"sync"
)

var baseUrl string = "https://avatars.githubusercontent.com/u/"
var folderAvatars string = "avatars"

func init() {
	err := os.MkdirAll(folderAvatars, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var wg sync.WaitGroup
	for id := 100000; id < 100100; id++ {
		wg.Add(1)
		go downloadAvatar(id, true, &wg)
	}
	wg.Wait()
}

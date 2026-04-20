package main

import "log"

// TODO: go mod init avatars
// TODO: go run .
const (
	URL = "https://avatars.githubusercontent.com/u/"
)

func main() {
	var id uint
	for id = 1; id < 50; id++ {
		result := getAvatar(id)
		log.Println(result)
	}
}

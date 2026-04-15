// https://avatars.githubusercontent.com/u/
// TODO: go mod init github
// TODO: go run .
package main

import (
	"log"
	"os"
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
	downloadAvatar(1, true)
	// TODO: Скачать 1000 аватаров
	// TODO: Измерить скорость выполнения
	
	// TODO: Переведите downloadAvatar в горутины
	// TODO: Измерить скорость выполнения
}

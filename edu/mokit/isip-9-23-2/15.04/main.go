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
	// TODO: Создать функцию downloadAvatar
	// TODO: параметры: code, id
	// TODO: Сохранения изображения в папку avatars
	// TODO: Название файла индентичное id
	avatar := getAvatar(1)
	downloadAvatar(avatar, 1)
}

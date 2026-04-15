// https://avatars.githubusercontent.com/u/
// TODO: go mod init github
// TODO: go run .
package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	err := os.Mkdir("avatars", 0644)
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
	fmt.Println(avatar)
}

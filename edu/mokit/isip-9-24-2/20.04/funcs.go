package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	os.MkdirAll("avatars", 0755)
}
// TODO: Создать функцию которая сохраняет
// TODO: изображение в папку avatars
// TODO: Функции: getAvatar, downloadAvatar, initAvatar
// TODO: initAvatar - запускает getAvatar и downloadAvatar


func getAvatar(id uint) []byte {
	url := fmt.Sprintf("%s%d", URL, id)
	response, err := http.Get(url)
	if err != nil {
		return []byte{}
	}
	if response.StatusCode != http.StatusOK {
		return []byte{}
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}
	}
	return data
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	url = "https://avatars.githubusercontent.com/u/"
)

func getByteImage(id uint) []byte {
	request := fmt.Sprintf("%s%d", url, id)
	response, err := http.Get(request)
	if err != nil {
		return []byte{}
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}
	}
	return data
}

func main() {
	var id uint
	for id = 1; id < 10; id++ {
		photo := getByteImage(id)
		fmt.Println(photo)
		// TODO: Создать БД
		// TODO: Создать таблицу
		// TODO: Таблица хранит 2 колонки
		// TODO: 1. id
		// TODO: 2. binary || byte
		// TODO: Все изображения хранить в БД
		// ? modernc.org/sqlite
		// ? database/sql
	}
}

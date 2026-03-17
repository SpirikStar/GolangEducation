package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Info struct {
	FirstName string `json:"first_name"`
	Year      uint16 `json:"year"`
}

func main() {
	data, err := os.ReadFile("info.json")
	if err != nil {
		fmt.Println("Error read file:", err)
		return
	}
	var info Info
	json.Unmarshal(data, &info)
	fmt.Println(info)
	info.Year += 10

	bytes, _ := json.Marshal(info)
	os.WriteFile("info.json", bytes, 0644)

	// TODO: создать функции:
	// TODO: 1) Функция создания json (2)
	// TODO: 2) Функция чтения json
	// TODO: 3) Функция += 1 (работа с count)
	// TODO: 4) Функция -= 1 (работа с count)
	// TODO: Файл json перезаписывать!
	// TODO: JSON:{"count":0, "is_zero":true}
	// TODO: true = is_zero == 0
	// TODO: false = is_zero not == 0
	// TODO: Не использовать пользовательский ввод
	// https://metanit.com/go/tutorial/7.1.php
	// https://metanit.com/go/tutorial/7.2.php
	// https://metanit.com/go/tutorial/7.3.php
}

package main

import (
	"encoding/json"

	"os"
)

type Config struct {
	HostName string `json:"host"`
	Port     int    `json:"port"`
	Debag    bool   `json:"debag"`
}

func main() {
	config := Config{
		HostName: "127.0.0.1",
		Port:     8000,
		Debag:    true,
	}
	bytes, _ := json.MarshalIndent(config, "", "    ")
	os.WriteFile("config.json", bytes, 0644)

	// content, _ := os.ReadFile("config.json")
	// var readContent Config
	// json.Unmarshal(content, &readContent)
	// fmt.Println(readContent)
	
	// TODO: Дополнить программу, чтобы она запрашивала
	// TODO: данные от пользователя
	// ! Каждая новая запись ОТДЕЛЬНЫЙ JSON!!!
}

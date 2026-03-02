package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Setting struct {
	Theme    string `json:"theme"`
	TextSize int    `json:"size"`
	TextBold bool   `json:"bold"`
}

func main() {
	setting := Setting{
		Theme:    "dark",
		TextSize: 14,
		TextBold: true,
	}
	byte_json, _ := json.MarshalIndent(
		setting,
		"",
		"    ",
	)
	os.WriteFile("setting.json", byte_json, 0644)

	// ? Чтение файла
	var settingTwo Setting
	loaded, _ := os.ReadFile("setting.json")
	json.Unmarshal(loaded, &settingTwo)
	fmt.Println(settingTwo)

}

// TODO: Самостоятельно создайте json файл
// ? {"count": 0}
// TODO: Создайте программу.
// TODO: 2 действия
// TODO: 1. Сбросить счёт в count (0)
// TODO: 2. Увеличить count += 1
// TODO: Результат сохранять в json

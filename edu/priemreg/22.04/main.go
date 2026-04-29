package main

import (
	"errors"
	"fmt"
	"os"
)

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("not divide zero")
	}
	return a / b, nil
}

func loadConfig(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("not file %s: %w", filename, err)
	}
	return data, nil
}
func main() {
	// ? Вывод: 0 not divide zero
	fmt.Println(divide(1, 0))
	
	// ? Вывод: 2 <nil>
	fmt.Println(divide(2, 1))


	fmt.Println(loadConfig("main.py"))

	// ? Вывод: [] not file config.json: open config.json: no such file or directory
	fmt.Println(loadConfig("config.json"))
}

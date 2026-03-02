package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	content := "bold text!"
	os.WriteFile("fileOne.txt", []byte(content), 0644)

	data, err := os.ReadFile("fileOne.txt")
	if err != nil {
		fmt.Println("Произошла ошибка чтения...")
		return
	}
	fmt.Println(data, string(data))

	file, _ := os.Open("info.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		fmt.Println(row)
	}

	var content_words strings.Builder
	words := []string{"is", "value", "key"}
	for _, value := range words {
		content_words.WriteString(fmt.Sprintf("\n%s", value))
	}
	os.WriteFile("words.txt", []byte(content_words.String()), 0644)

	// TODO: Создайтеи3 функцию которая принимает слово
	// TODO: и ищет его в words.txt
	// TODO: Функция должна вернуть true || false
}

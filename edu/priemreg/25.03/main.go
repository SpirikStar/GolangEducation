package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	entries, _ := os.ReadDir(".")

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		content := readFile(entry.Name())
		if len(content) == 0 {
			continue
		}
		searchContent(content, "hello")
	}
}
func searchContent(content []byte, value string) {
	text := string(content)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), strings.ToLower(value)) {
			fmt.Println("Успех")
		}
	}
}

func readFile(filename string) []byte {
	content, _ := os.ReadFile(filename)
	return content
}
// TODO: Если поиск успешный, необходимо чтобы searchContent сообщил в каком файле
// TODO: Вывод:
// TODO: (file1.txt) Поиск 1/3. Не найдено. 
// TODO: (file2.txt) Поиск 2/3. Успешно!
// TODO: (file3.txt) Поиск 3/3. Не найдено. 
// TODO: Создайте 3 worker функции, которые используют searchContent
// TODO: Названия файлов необходимо записывать в буфер. Буфер ограничен 5 элементах
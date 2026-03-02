package main

import "fmt"

func addWord(array *[]string, value any) {
	// TODO: Ваша задача value добавить в массив words
	// TODO: Не добавлять в массив значения которые не являются string
}

func searchWord(array []string, value any) string {
	// TODO: Ваша задача найти слово в массиве
	// TODO: 2 результата: "Результат отсутсвует" или "Успех: ..."
	return "Результат отсутсвует"
}

func main() {
	words := []string{}
	addWord(&words, "hello")
	result := searchWord(words, "h")
	fmt.Println(result)
}

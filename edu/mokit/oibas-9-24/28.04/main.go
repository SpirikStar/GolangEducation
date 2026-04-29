package main

import "fmt"

func operation(nOne, nTwo float32) (float32, error) {
	if nTwo == 0 {
		return 0, fmt.Errorf("E101: INFO: %f / %f", nOne, nTwo)
	}
	return nOne / nTwo, nil
}
func elems(v ...any) {}
func main() {
	words := []string{} // ! Не перемещать его
	// TODO: создать функцию которая добавляет новое слово в words
	// TODO: Подсказка: ..., any, &, *
	// TODO: Без дубликатов; только срока; хотя бы один аргумент
	// TODO: Слово без пустоты - пример: ""
	elems(1, 2, 3, 4, "popopop")
	fmt.Println(operation(10, 0))
	fmt.Println(operation(10, 10))
}

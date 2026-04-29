package main

import (
	"fmt"
)

func operation(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("error")
	}
	return a / b, nil
}
// func point (elems ...string) {}
func main() {
	// TODO: Задание. 
	context := []string{} // ! Не перемещать
	// TODO: Создать функцию которая принимает (∞ аргументы ...)
	// TODO: Функция должна заполнить context
	// TODO: В context не может сохраняться пустота
	// TODO: В context не может сохраняться всё кроме string
	// TODO: Логировать ошибки в debag.log (см. 06.04)
	// TODO: Подсказки: ..., *, &	
	// TODO: В context не должны бьть дубликаты
	// ! Не использовать готовые инструменты по поиску
	// ? До 16:40 - решение
	// TODO: Создать функцию которая чистит context
	// ! Не использовать готовые инструменты по чистке
	// ? Требования (оценка 3): ответ на 3 вопроса по коду
	// ? Требования (оценка 2): готовый код
	// ? Итог: 3+2 = оценка 5

	fmt.Println(operation(1, 0.5))
}

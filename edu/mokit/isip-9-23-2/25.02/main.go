package main

import "fmt"

type Messsage struct {
	Id      int
	Content string
}

// Способ 1: Метод копии (редко)
func (m Messsage) returnMessage() {
	fmt.Println(m.Content)
}

// Способ 2: Метод приёмника (указатель)
func (m *Messsage) returnId() {
	m.Id++
	fmt.Println(m.Id)
}

func funcOne(elems ...int) (bool, int) {
	if len(elems) > 0 {
		return true, 1
	}
	return false, 0
}

func main() {
	message := Messsage{1, "Hello!"}
	message.returnMessage()
	funcOne(1, 2, 3, 4)
	funcOne()
}

Создать функцию, которая принимает аргументы.
Функция должна возращать 3 значения
1. Статус (если аргументы не переданы false)
2. Если статус == true вернуть максимальный аргумент
3. Если статус == true найти среднее значение

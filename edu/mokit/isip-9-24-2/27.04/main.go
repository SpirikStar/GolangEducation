package main

import "fmt"

func division(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Error: b == %f", b)
	}
	return a / b, nil
}
// ? func test(elems ...string){

// ? }
// ? test("1", "2", "3")
func main() {
	messages := []string{}
	// TODO: Cоздать функцию которая добавляет сообщение в message
	// TODO: Пустое сообщение нельзя сохранить в message
	// TODO: Числа нельзя передавать в функцию
	// TODO: Сообщения можно передавать несколько (...)
	// TODO: * Все действия записывать в log файл
	// TODO: Функция возращает 2 результата.
	fmt.Println(division(1.2, 0))
}

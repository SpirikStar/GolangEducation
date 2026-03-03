package main

import "fmt"

func funcOne() {
	fmt.Println("start (1)")
}
func funcTwo(a int, b int) {
	fmt.Println(a + b)
}
func funcThree(a, b int) {
	fmt.Println(a + b)
}
func funcFour(elem any) {
	fmt.Println(elem)
}

func funcActive(status bool) {
	if !status {
		return
	}
	fmt.Println("Access funcs (2)")
}

// ! https://iqnix.link/func84
// TODO: Дедлайн - 15:17
// TODO: Ответить на 1 вопрос
// ! ДО "Неопределенное количество параметров"

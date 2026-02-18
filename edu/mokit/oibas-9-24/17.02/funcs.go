package main

import "fmt"

func funcOne() {
	fmt.Println("Функция сработала (1)")
}

func funcTwo(a int, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func funcThree(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func funcFour(elem any) {
	switch value := elem.(type) {
	case int:
		fmt.Printf("Целое число: %d\n", value)
	case string:
		fmt.Printf("Строка: %s\n", value)
	default:
		fmt.Printf("Неизвестно: %T\n", value)
	}
}

func funcFive(addr *int) {
	*addr += 1
}

func funcAppend(array *[]int, elem any) {
	switch value := elem.(type) {
	case int:
		*array = append(*array, value)
	}
}

func funcSix() int {
	return 100
}
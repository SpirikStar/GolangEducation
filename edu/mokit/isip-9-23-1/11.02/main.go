package main

import (
	"fmt"
	"time"
)

func main() {
	funcOne()
	funcTwo("is", "words")

	count := 0
	funcThree(&count)
	fmt.Println(count)
	
	dTime := time.Now()
	dTimeUnix := time.Now().UnixNano()
	fmt.Println(dTime, dTimeUnix)
	// TODO: Самостоятельно изучить библиотеку
	// ! "math/rand" "time"
	// TODO: Создайте функцию, которая принимает
	// TODO: от пользователя N - число

	// TODO: Функция должна сгенерировать код доступа
	// TODO: Код должен состоять из длины N чисел!
	// TODO: Функция должна вернуть код
}

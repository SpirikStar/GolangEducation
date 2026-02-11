package main

import (
	"fmt"
)

func main() {
	count := 0
	for count < 5 {
		count += 1
	}

	for {
		if count > 0 && count < 10 {
			break
		}
		count++
	}
	number := 100
loop:
	for number < 120 {
		switch number {
		case 105:
			break loop
		default:
			number++
			fmt.Printf("Вывод: %d\n", number)
		}
	}

	numberTwo := 1
	for numberTwo < 10 {
		numberTwo++

		if numberTwo%2 == 0 {
			fmt.Println(numberTwo)
			continue
		}

		fmt.Println("----")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	message := "is home"
	for index, value := range message {
		if value == ' ' {
			continue
		}
		fmt.Printf(
			"index: %d; value: %s \n",
			index,
			string(value),
		)
	}

	textOne := "hello"
	textTwo := "two"
	res := textOne + textTwo
	res += "!"
	fmt.Println(res)

	var userMessage string
	fmt.Scan(&userMessage)

	// TODO: Что такое for
	// TODO: Отличия continue и break
	
	// TODO: Алгоритм работы:
	// ? Введите сообщение: is home!
	// ? Итог сообщения: ishome
	// ? Введите сообщение: is ivan green!
	// ? Итог сообщения: isivan
	// ! max len(6)
}

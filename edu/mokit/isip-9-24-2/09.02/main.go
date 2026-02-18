package main

import (
	"fmt"
)

func main() {
	count := 0
	for count < 4 {
		count++
	}

	for {
		if count > 0 && count < 10 {
			break
		}
		count++
	}
	numberOne := 0
	for numberOne < 10 {
		numberOne++
		if numberOne%2 == 0 {
			continue
		}
		fmt.Printf("Вывод: %d \n", numberOne)
	}
	for i := 0; i < 15; i++ {
		fmt.Printf("Вывод: %d \n", i)
	}

	message := "is home"
	for index, value := range message {
		if value == ' ' {
			continue
		}
		fmt.Printf("index: %d; value: %s \n", index, string(value))
	}
	// go run main.go
	var userMessage string
	fmt.Print("Введите сообщение: ")
	fmt.Scan(&userMessage)
}

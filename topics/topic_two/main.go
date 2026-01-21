package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	count := 0
	for count < 10 {
		count++
	}
	fmt.Println("Result:", count)

	for {
		var userKey string
		fmt.Print("Введите команду: ")
		fmt.Scan(&userKey)
		if userKey == "exit" {
			fmt.Println("Закрытие программы...")
			break
		}
	}


	text := "Hello"
	for index, char := range text {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}
}

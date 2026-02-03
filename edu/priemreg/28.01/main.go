package main

import "fmt"

func main() {
	number := 0

	for number < 5 {
		number++
	}
	fmt.Printf("Number: %d\n", number)

	for {
		if number == 15 {
			fmt.Printf("Number: %d\n", number)
			break
		}
		number++
	}

loop: // метка
	for {
		switch number {
		case 30:
			fmt.Printf("Number: %d\n", number)
			break loop // обращение к метке
		default:
			number++
		}
	}
	newNumber := 0
	for i := 0; i < 5; i++ {
		newNumber += i
	}
	fmt.Println(newNumber)

	text := "hello"
	for index, value := range text {
		if value == 'l' {
			fmt.Println("Success")
		}
		fmt.Printf("index: %d | value: %s\n", index, string(value))
	}

	email := "bobik@ya.ru"
	isEmail := 0
	for _, value := range email {
		if value == '@' {
			isEmail += 1
		}
	}
	if isEmail == 1 {
		fmt.Println("Почта верная!")
	} else {
		fmt.Println("Ошибка в почте!")
	}

	var a, b, c int
	fmt.Println(a, b, c)

	// Массивы.

	// value: [0, 10, 0]
	// index: [0, 1, 2]
	var nums [3]int
	nums[1] = 10
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(nums[2])

	fmt.Println("Вывод массива:")
	for index, value := range nums {
		fmt.Printf("index: %d | value: %d\n", index, value)
	}
}

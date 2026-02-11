package main

import "fmt"

func main() {
	number := 150

	// TODO: && - И
	// ? false && true = false
	// TODO: || - ИЛИ
	// ? false && true = true

	if number > 0 && number < 100 {
		fmt.Printf("if: 100 > %d > 0\n", number)
	} else if number == 100 {
		fmt.Printf("if: 100 == %d\n", number)
	} else if number == 120 {
		fmt.Printf("if: 100 == %d\n", number)
	} else {
		fmt.Println("if: Исключение")
	}

	switch number {
	case 10:
		fmt.Printf("switch: 10 == %d\n", number)
	case 20:
		fmt.Printf("switch: 20 == %d\n", number)
	default:
		fmt.Println("switch: Исключение")
	}

	switch {
	case number < 0:
		fmt.Printf("switch 2: %d < 0\n", number)
	}

	var userVar int
	fmt.Print("Введите значение: ")
	fmt.Scan(&userVar)

	a, b, c := 1, 1.1, "text"
	fmt.Println(a, b, c)

	var userInt int
	fmt.Println(userInt)
}

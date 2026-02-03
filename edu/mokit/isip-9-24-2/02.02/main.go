package main

import "fmt"

func main() {
	number := 100

	// TODO: && - И
	// TODO: || - ИЛИ
	// TODO: true && false = false
	// TODO: true || false = true
	if number > 0 && number < 100 {
		fmt.Printf("0 < %d < 100\n", number)
	} else if number == 100 {
		fmt.Printf("%d == 100\n", number)
	} else if number == 150 || number == 200 {
		fmt.Printf("%d == 150\n", number)
	} else {
		fmt.Println("if: Сработало исключение!")
	}

	switch number {
	case 200:
		fmt.Printf("switch: %d == 200\n", number)
	case 400:
		fmt.Printf("switch: %d == 400\n", number)
	case 600:
		fmt.Printf("switch: %d == 600\n", number)
	default:
		fmt.Println("switch: Сработало исключение!")
	}

	switch {
	case number > 1:
		fmt.Printf("switch 2: %d > 1\n", number)
	}

	var userElem int
	fmt.Print("Введите число: ")
	fmt.Scan(&userElem)

	a, b, c := 1, 2.9, 3
	var u int // default: 0
	fmt.Println(a, b, c, u)
}

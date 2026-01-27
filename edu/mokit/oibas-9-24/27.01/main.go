package main

import "fmt"

func main() {
	number := 0
	copyNumber := number
	copyNumber++

	linkNumber := &number
	*linkNumber += 15
	fmt.Println(number)
	// TODO: true and true = true
	// TODO: false and true = false
	if number > 0 && number < 10 {
		fmt.Printf("Переменная: %d > 0 и %d < 10\n", number, number)
	} else if number >= 10 && number == 12 || number > 15 {
		fmt.Printf("Переменная: %d > 0 и %d < 15\n", number, number)
	} else if number == 16 {
		fmt.Printf("Переменная: %d == 16\n", number)
	} else {
		fmt.Println("Except!")
	}

	var count int = 0
	switch count {
	case 1:
		fmt.Printf("Value: %d == 1", count)
	case 2:
		fmt.Printf("Value: %d == 2", count)
	default:
		fmt.Println("Value - Except!")
	}

	switch {
	case count > 2 && count < 5:
		fmt.Printf("Result: 5 > %d > 2\n", number)
	case count > 2 && count < 5:
		fmt.Printf("Result: 5 > %d > 2\n", number)
	}
}

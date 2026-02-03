package main

import "fmt"

func main() {
	number := 0
	for number < 5 {
		fmt.Printf("current number: %d\n", number)
		number++
	}
	for {
		if number < 10 {
			number++
			continue
		}
		fmt.Printf("result number: %d\n", number)
		break
	}
loop:
	for {
		switch number {
		case 20:
			fmt.Printf("result switch: %d\n", number)
			break loop
		default:
			number++
		}
	}

	text := "hello"
	for index, value := range text {
		fmt.Printf(
			"index: %d | value: %s\n",
			index, string(value),
		)
	}

	var message string
	fmt.Print("Введите сообщение: ")
	fmt.Scan(&message)

	var chars, space, letters int
	for _, value := range message {
		if value == '.' || value == ',' {
			chars++
		} else if value == ' ' {
			space++
		} else {
			letters++
		}
	}

	for n := 0; n < 5; n++ {
		fmt.Printf("Value: %d \n", n)
	}

	
}

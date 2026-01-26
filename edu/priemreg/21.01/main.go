package main

import "fmt"

func main() {
	number := 10
	numberTwo := number

	numberTwo += 1 // 10 + 1 = 11
	numberTwo++    // 11 + 1 = 12

	fmt.Println(number, numberTwo, &number)

	maskNumber := &number
	*maskNumber++
	fmt.Println(number, *maskNumber)

	isOne := true  // 1
	isTwo := false // 0

	if isOne {
		fmt.Println("Правда")
	}
	if isTwo {
		fmt.Println("Ложь")
	}
	// >=; <=; >; <; ==
	// && - И
	// || - ИЛИ
	// true && true = true
	// true && false  = false

	// true || true = true
	// true || false = true
	// false || false = false
	userNumber := 210
	if userNumber > 0 && userNumber < 100 {
		fmt.Printf("100 > %d > 0 \n", userNumber)
	} else if userNumber > 100 && userNumber < 150 {
		fmt.Printf("150 > %d > 100 \n", userNumber)
	} else if userNumber < 200 || userNumber == 180 {
		fmt.Printf("%d < 200 \n", userNumber)
	} else {
		fmt.Println("Нет нужного выражения")
	}

	balance := 200.0
	switch balance {
	case 100.0:
		fmt.Printf("%.1f == 100.0\n", balance)
	case 200.0:
		fmt.Printf("%.1f == 200.0\n", balance)
	default:
		fmt.Println("Выражения switch-case => нет нужного!")
	}

	switch {
	case balance > 100.0:
		fmt.Printf("%.1f > 100.0\n", balance)
	case balance > 200.0:
		fmt.Printf("%.1f > 200.0\n", balance)
	}

	var usNum int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&usNum)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	result := usNum*2 + 10
	fmt.Printf("Result: %d\n", result)
}

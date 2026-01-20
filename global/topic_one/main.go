package main

import "fmt"

func main() {
	numberOne := 10
	if numberOne == 10 {
		fmt.Println("numberOne == 10")
	}

	numberTwo := 20
	if numberTwo < 10 {
		fmt.Println("numberTwo < 10")
	} else if numberTwo < 20 {
		fmt.Println("numberTwo < 20")
	} else {
		fmt.Println("numberTwo >= 20")
	}

	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	if age > 0 && age < 18 {
		fmt.Println("Возраст в диапозоне от 0 до 18")
	}

	switch age {
	case 14:
		fmt.Println("Пора получать паспорт")
	case 21:
		fmt.Println("Пора менять паспорт")
	default:
		fmt.Println("Исключение")
	}

	switch {
	case age > 18:
		fmt.Println("Вы взрослый!")
	}
}

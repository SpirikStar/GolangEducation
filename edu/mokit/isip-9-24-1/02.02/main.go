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

	// TODO: int = 0
	// TODO: float = 0.0
	// TODO: string = ""
	var userInt int
	fmt.Println(userInt)

// TODO: Срок проверки до 15:25
// TODO: Запросить у пользователя число
// TODO: Число должно быть от 0 до 100
// TODO: Вывод если от 0 до 25 = 2
// TODO: Вывод если от 25 - 50 = 3
// TODO: Вывод если от 50 - 75 = 4
// TODO: Вывод если от 75 - 100 = 5
// TODO: Если число отрицательное сообщить об этом

// TODO: оценка 3:	 код + 1 вопрос по коду + 1 вопрос по лекции
// TODO: оценка 4: код + 1 вопрос по коду + 2 вопроса по лекции
// TODO: оценка 5: код + 1 вопрос по коду + 3 вопроса по лекции
}

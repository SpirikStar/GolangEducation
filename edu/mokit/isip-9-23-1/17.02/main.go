package main

import "fmt"

type Person struct {
	FirstName string
	Age       int
	Year      int
}

func main() {
	userOne := Person{"Ivan", 7, 2019}
	userTwo := Person{"Maria", 12, 2014}
	userOne.Age++
	userOne.Year--
	fmt.Println(userOne, userOne.FirstName)
	fmt.Println(userTwo)

	var userThree Person
	userThree.FirstName = "Bobik"
	userThree.Age = 3
	userThree.Year = 2023
	fmt.Println(userThree)

	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2))

	func(message string, p Person) {
		fmt.Println(message, p.FirstName)
	}("start in func...", userTwo)

	// TODO: Калькулятор скидок товаров

	type Product struct {
		// ? Создать поля: Name, Price
	}
	// TODO: Создать массив из 3 товаров
	// TODO: Создать анонимную функцию, которая
	// TODO: высчитывает скидку товаров
	// TODO: на основе цены
}

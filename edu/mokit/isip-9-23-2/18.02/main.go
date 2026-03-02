package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	FirstName string
	Age       int
	Active    bool
}

func main() {
	userOne := Person{"Bobik", 3, true}
	userOne.Age += 1
	userTwo := Person{"Ivan", 12, false}
	userTwo.Active = true
	fmt.Println(userOne, userTwo)
	fmt.Println(userOne.Age)

	var userThree Person
	userThree.FirstName = "Gleb"
	userThree.Age = 10
	fmt.Println(userThree)

	numOne := func() int {
		return 100
	}
	numTwo := func(a, b int) int {
		return a + b
	}
	func() {
		fmt.Println("Start func!")
	}()
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)
	fmt.Println(numOne(), numTwo(1, 2))

	file, _ := os.Open("file.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// TODO: Создайте файл words.txt
	// TODO: Вручную поместите 10-15 слов
	// TODO: Создайте функцию search_words
	// TODO: Создайте параметр value any
	// TODO: Алгоритм работы:
	// ? result := search_words('Бо')
	// ? result := ["Боба", "Большой"]
	// TODO: лимит слов в массиве 5
	// ! "bufio" "fmt" "os" "string"
	// ! Требования:
	// ! 1. Выдержать 3 атаки
	// ! 2. Ответить на 3 вопроса по лекции 
	// ! 3. Ответить на вопросы по коду
}

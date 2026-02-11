package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	defer fmt.Println("Процессы завершены!")
	defer funcMessage("Ivan")

	funcStart()
	funcMessage("Bobik")
	add(1, 2)

	joinResult := joinWords("is", "home")
	fmt.Println(joinResult)

	var varOne any = 10
	varOne = "text"
	varOne = 10.4
	fmt.Println(varOne)

	// TODO: Как пробижаться по массиву?

	arrayOne := [...]any{1, "text", 1.1}
	for _, value := range arrayOne {
		fmt.Printf("type: %T; value: %v \n", value, value)
	}

	mapInfo := map[string]any{
		"first_name": "Bobik",
		"age":        3,
	}
	fmt.Println(arrayOne, mapInfo)

	count := 0
	fmt.Println("До: ", count)
	increment(&count)
	fmt.Println("После: ", count)

	dTime := time.Now()
	dTimeUnix := dTime.UnixNano()
	fmt.Println(dTime, dTimeUnix)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println(rnd.Intn(100))
	fmt.Println(rnd.Float64())

	// TODO: Создать новый проект
	// TODO: Создать 2 файла
	// TODO: 1 файл - main.go
	// TODO: 2 файл - funcs.go
	// TODO: Создать функцию для генерации пароля
	// ! Использовать только: "fmt" "math/rand" "time"
	// TODO: Функция принимает 2 параметра
	// TODO: 1 параметр - len пароля
	// TODO: 2 параметр - включать ли в пароль числа
	// TODO: true: текст + числа; false: текст
	// TODO: Функция должна вернуть пароль
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	// var имя [длина]тип

	var numbers [5]int
	fmt.Println("массив из 5 целых чисел (по умолчанию заполнен нулями):", numbers)
	names := [2]string{"Ivan", "Bobik"}
	fmt.Println("массив из 2 строк:", names)

	elems := [...]int{10, 20, 30}
	elems[1] += 20
	dubl_elems := elems
	dubl_elems[0] *= 2
	fmt.Println("Дубликат elems:", dubl_elems)
	fmt.Println("компилятор сам определит длину:", elems)

	fmt.Println("Обращение по индексу:", elems[0])
	fmt.Println("Длина массива:", len(elems))

	colors := [3]string{"red", "green", "black"}
	for index, value := range colors {
		fmt.Printf("Index %d: %s\n", index, value)
	}

	fmt.Println("---Срезы---")

	elemsOne := []int{30, 40, 50}
	elemsOne = append(elemsOne, 60)
	elemsOne = append(elemsOne, 60, 70, 80)
	fmt.Println("Срез:", elemsOne)

	elemsTwo := make([]int, 3)
	fmt.Println("длина = 3, ёмкость = 3:", elemsTwo)
	
	elemsThree := make([]int, 3, 5)
	fmt.Println("длина = 3, ёмкость = 5:", elemsThree)
	fmt.Println("длина =", len(elemsThree))
	fmt.Println("ёмкость =", cap(elemsThree))


	fmt.Println("Поиск элемента elemsOne:", slices.Contains(elemsOne, 50))
	fmt.Println("Поиск элемента elemsThree:", slices.Contains(elemsThree, 50))
}

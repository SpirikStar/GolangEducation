package main

import (
	"fmt"
)

func main() {
	//? Формула создания массива: var имя [N]тип

	var arrOne [3]int      // TODO: default int = 0
	fmt.Println(arrOne)    // TODO: вывод[0, 0, 0]
	fmt.Println(arrOne[0]) // TODO: indexs: 0,1,2

	arrOne[1] = 10
	fmt.Println(arrOne) // TODO: вывод[0, 10, 0]
	arrOne[1] += 15
	fmt.Println(arrOne) // TODO: вывод[0, 25, 0]

	//? Определения массива: имя := [...]тип{}
	//? ... - auto определения len массива
	//? ... - использовать можно только с {}
	days := [...]string{"Пн", "Вт", "Ср"}
	fmt.Println(days) // TODO: вывод["Пн","Вт","Ср"]

	for index, value := range days {
		fmt.Printf(
			"index: %d; value: %s \n",
			index,
			value,
		)
	}

	var count int
	for count < 5 { // TODO: пока true {действие}
		count++ // TODO: count += 1
	}
	fmt.Println(count) // TODO: вывод 5

	arrTwo := []int{1, 2, 3} 
	arrTwo = append(arrTwo, 4)
	fmt.Println(arrTwo)
}

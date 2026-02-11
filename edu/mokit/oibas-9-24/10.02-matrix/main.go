package main

import "fmt"

func main() {
	var matrixOne [3][5]int
	fmt.Println("Вложенный массив (all default 0): ", matrixOne)

	matrixTwo := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Вложенный массив: ", matrixTwo)

	matrixPartial := [3][3]int{
		{1, 2},
		{3},
	}
	matrixPartial[1][1] = 4
	fmt.Println("Вложенный массив (partial): ", matrixPartial)

	matrixOriginal := [2][1]int{}
	matrixCopy := matrixOriginal
	matrixCopy[1][0] = 5
	fmt.Println("Вложенный массив: ", matrixOriginal)

	var numbers [3][3]int
	for _, row := range numbers {
		fmt.Println(row)
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i][i])
	}
}

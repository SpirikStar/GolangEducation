package main

import "fmt"

func main() {
	arrOne := [3]int{10}
	arrOne[0] += 5
	fmt.Println("Array (1): ", arrOne, len(arrOne))
	for index, value := range arrOne {
		fmt.Printf("i: %d; v: %d\n", index, value)
	}

	arrTwo := [2]any{}
	arrTwo[0] = "text"
	arrTwo[1] = 100
	fmt.Println("Array (2): ", arrTwo)

	arrThree := [...]string{"is", "home"}
	fmt.Println("Array (3): ", arrThree, len(arrThree))
	
	
	arrFour := []int{}
	arrFour = append(arrFour, 1)
	arrFour = append(arrFour, 2, 3)
	fmt.Println("Array (4): ", arrFour)


	arrFive := [10]int{}
	// TODO: Используйте цикл for!
	// TODO: Итог: [1 0 2 0 3 0 4 0 5 0]
}

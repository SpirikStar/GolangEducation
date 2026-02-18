package main

import "fmt"

func main() {
	arrOne := [3]int{5}
	arrOne[0] += 10
	fmt.Println("Array (1): ", arrOne)

	for index, value := range arrOne {
		fmt.Printf("i: %d; v: %d\n", index, value)
	}

	arrTwo := [2]any{}
	arrTwo[0] = 100
	arrTwo[1] = 100.4
	fmt.Println("Array (2): ", arrTwo)

	for _, value := range arrTwo {
		fmt.Printf("type: %T - value: %v", value, value)
	}

	arrThree := [...]string{"is", "home"}
	fmt.Println("Array (3): ", arrThree)

	arrFour := []int{}
	arrFour = append(arrFour, 1, 2)
	fmt.Println("Array (4): ", arrFour, len(arrFour))

	history := [10]any{"is", "help", 2}
	result := [8]int{}

	for index, value := range history {
		// TODO: отфильтровать history по числам.
		// TODO: Итог поместить в result
	}
	// TODO: result = [1 0 2 0 3 0 4 0]
}

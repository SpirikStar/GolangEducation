package main

import "fmt"

func calc(a, b int) int {
	return a + b
}

func calcUpgrade(nums ...int) {
	total := 0
	for _, value := range nums {
		total += value
	}
	fmt.Println(total)
}

func main() {
	fmt.Println(calc(1, 2))

	result := func() int {
		return 100
	}
	resCalc := func(a, b int) int {
		return a + b
	}
	fmt.Println(result(), resCalc(1, 2))
	func() {
		fmt.Println("Event func (1)")
	}()

	calc(func() int { return 100 }(), 2)

	calcUpgrade()
	calcUpgrade(1, 1, 1)

	arrayOne := []int{1, 2}
	// Извлечение данных из массива
	calcUpgrade(arrayOne...)

	var a any = "100"
	val, status := a.(int)
	if status {
		fmt.Println(val)
	}
}

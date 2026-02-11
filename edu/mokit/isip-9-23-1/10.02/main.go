package main

import "fmt"

func Log() {
	fmt.Println("Логирование программы...")
}

func Add(a, b int) int {
	Log()
	return a + b
}

func FuncParams(elems []int) {
	fmt.Println("Массив (1): ", elems)
}

func PrintValue(elem any){
	fmt.Printf("Тип: %T, Значение: %v\n", elem, elem)
}

func main() {
	PrintValue(100.3)
	var a any
	a = 40
	a = "text"
	a = []int{1, 2, 3}

	var b interface{}
	fmt.Println("Псеводонимы: ", a, b)


	elems := []int{1, 2, 3}
	FuncParams(elems)

	defer fmt.Println("Конец: ", 1)
	defer fmt.Println("Конец: ", 2)
	defer Log()

	fmt.Println(Add(1, 2))
}

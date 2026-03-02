package main

import "fmt"

func funcOne() {
	fmt.Println("Запуск...")
}

func funcPrint(value any) {
	fmt.Println(value)
}

func funcThree(a int, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func funcFour(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func funcFive(elem int) {
	if elem < 0 {
		return
	}
	fmt.Println("Current number > 0")
}

func funcSix(a, b int) int {
	return a + b
}

func addCount(addr *int) {
	*addr += 5
}
func main() {	
	result := funcSix(1, 2)
	fmt.Println(result)
	funcOne()
	funcPrint("is home")
	funcPrint("is value")
	funcPrint(100)
	funcThree(1, 2)

	var elem any = 100

	_, ok := elem.(int)
	if ok {
		fmt.Println("Данный элемент является int")
	}

	funcFive(-10)
	funcFive(10)

	count := 0
	addCount(&count)
	addCount(&count)
	addCount(&count)
	fmt.Println(count)

}

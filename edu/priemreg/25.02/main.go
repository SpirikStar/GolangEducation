package main

import "fmt"

func funcNums(elems ...int) {
	fmt.Println(elems)
}
func funcReturns(elems ...int) (bool, string) {
	if len(elems) > 0 {
		return true, "Массив содержит данные" 
	}
	return false, "Массив пустой" 
}
func main() {
	status, message := funcReturns(1, 2, 3)
	fmt.Println(status, message)
	
	func() {
		funcNums(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		funcNums()
		nums := []int{100, 200}
		funcNums(nums...)
	}()

	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)

	result := func(a, b int) int {
		return a + b
	}
	fmt.Println(result(1, 2))
	fmt.Println(result(2, 3))
}

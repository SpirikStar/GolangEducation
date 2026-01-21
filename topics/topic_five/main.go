package main

import "fmt"

func actionOne() {
	fmt.Println("Первая функция")
}
func actionTwo() int {
	return 100
}
func actionThree() float32 {
	return 10.2
}
func actionFour(a string, b int) {
	fmt.Println(a, b)
}
func actionSix(elem *int) {
	*elem += 1
}
func main() {
	actionOne()
	varFunc := actionTwo()
	fmt.Println(varFunc)
	actionFour("text", 2)

	varElem1 := 0
	actionSix(&varElem1)
	fmt.Println(varElem1)

}

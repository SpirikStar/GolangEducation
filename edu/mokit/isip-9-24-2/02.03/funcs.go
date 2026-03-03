package main

import "fmt"

func funcOne() {
	fmt.Println("Work func (1)")
}

func funcTwo(a int, b int) {
	fmt.Println("func (2):", a+b)
}

func funcThree(a, b int) {
	fmt.Println("func (3):", a+b)
}

func funcFour(status bool) {
	if !status {
		return
	}
	fmt.Println("func (4): success")
}

func funcSix() int {
	return 100
}

func funcAny(elem any) any {
	return elem
}

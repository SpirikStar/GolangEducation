package main

import (
	"fmt"
	"log"
)

func funcOne() (int, bool) {
	return 100, true
}

func funcTwo(nums ...int) {
	fmt.Println(nums)
}

func main() {
	val, isVal := func() (string, bool) {
		return "text", true
	}()
	fmt.Println(val, isVal)
	value, status := funcOne()
	fmt.Println(value, status)
	fmt.Println(funcOne())
	funcTwo()
	funcTwo(1, 2, 3)
	log.Fatalln	("Ошибка!")
}

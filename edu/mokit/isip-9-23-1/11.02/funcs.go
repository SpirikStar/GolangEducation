// ! go mod init lesson
// ! go run .
package main

import "fmt"

func funcOne() {
	fmt.Println("Функция сработала...")
}

func funcTwo(wordOne, wordTwo string) {
	fmt.Println(wordOne + " " + wordTwo)
}

func funcThree(elem *int) {
	*elem++
}

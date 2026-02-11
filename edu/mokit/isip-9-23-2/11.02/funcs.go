// ! go mod init lesson
// ! go run .
package main

import "fmt"

func funcStart() {
	fmt.Println("Сработала функция (1);")
}
func funcMessage(name string) {
	fmt.Printf("Hello, %s! \n", name)
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d \n", a, b, a+b)
}

func joinWords(word_one, word_two string) string {
	return word_one + " " + word_two
}

func increment(adr *int) {
	*adr++
}

// go run main.go funcs.go
// go mod init lesson
// go run .
package main

import "fmt"

func main() {
	funcOne()
	funcTwo(1, 2)

	funcFour(1)
	funcFour("is user?")

	count := 80
	funcFive(&count)
	funcFive(&count)
	fmt.Println(count)

	numbers := []int{}
	funcAppend(&numbers, 10)

	res := funcSix()
	fmt.Println("return from func - value: ", res)
}

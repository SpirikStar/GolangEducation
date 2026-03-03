// TODO: go run main.go funcs.go
// TODO: go mod init lesson
// TODO: go run .
package main

import "fmt"

func main() {
	funcOne()
	funcOne()
	funcTwo(1, 2)
	funcFour(false)
	result := funcSix()
	fmt.Println(result)
	fmt.Println(funcAny(1))
	fmt.Println(funcAny(1.2))
}

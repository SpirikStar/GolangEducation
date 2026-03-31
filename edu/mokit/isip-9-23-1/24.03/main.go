package main

import "fmt"

func main() {
	chOne := make(chan int, 3)
	chOne <- 10
	chOne <- 20
	chOne <- 20
	// chOne <- 20 // ! Ошибка. Нельзя превышать лимит

	fmt.Println(<-chOne) // clear
	fmt.Println(<-chOne) // clear
	fmt.Println(<-chOne) // bufer make null
}

package main

import "fmt"

func main() {
	arrOne := [...]int{1, 2, 3}
	fmt.Println(arrOne)

	arrTwo := []int{1, 2, 3}
	arrTwo = append(arrTwo, 4)
	fmt.Println("Способ «Литерал»:", arrTwo)

	arrThree := make([]int, 2)
	arrThree[0] += 10
	fmt.Println("Способ «make»:", arrThree, len(arrThree))

	var arrFour []string = make([]string, 0, 2)
	arrFour = append(arrFour, "Bobik")
	fmt.Println("Способ «make»:", arrFour, len(arrFour), cap(arrFour))
}

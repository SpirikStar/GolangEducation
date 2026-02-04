package main

import "fmt"

func main() {
	var arrOne [3]int // TODO: default [0, 0, 0]
	arrTwo := [3]int{1, 2, 3}
	arrThree := [...]int{1, 2, 3} // TODO: auto len

	fmt.Println(arrOne, arrTwo, arrThree)
	s1 := []int{1, 2}     // ? способ: литерал
	s1 = append(s1, 3, 4) // ! перезапись s1
	fmt.Println(s1)

	// ! Формула: make(типДанных(true), длина(true), ёмкость(false))
	s2 := make([]string, 2, 5) // ? способ: make
	s2 = append(s2, "one")
	fmt.Println(s2, s2[0], len(s2), cap(s2))
}

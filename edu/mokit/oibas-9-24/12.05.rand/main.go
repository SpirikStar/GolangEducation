package main

import (
	"log"
	"math/rand/v2"
)

type RecNum struct {
	nums []int
}

func main() {
	nums := []int{}
	// TODO: 1. Создать функцию, которая будет добавлять
	// TODO: рандомное число в этот массив.
	// TODO: 1.1. addNum(100, 1000, &nums) -> (кол-во, диапозон, адрес)
	// TODO: 2. Создать рекурсию, которая вернёт кол-во дубликатов
	// TODO: 2.1 searchDubl(15, 0, &nums) -> (что ищем, index, адрес)
	// TODO: Результат: кол-во дубликатов
	log.Println(rand.IntN(100))
}

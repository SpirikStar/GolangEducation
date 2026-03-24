package main

import "fmt"

type Item struct {
	Name string
	Code uint16
}

func main() {
	itemOne := Item{}
	itemOne.Name = "Ivan"
	itemTwo := Item{"Bob", 101}

	var itemThree Item = Item{
		Code: 102,
	}

	items := []Item{
		{Name: "Item 1", Code: 1000},
		{Name: "Item 2", Code: 1002},
	}
	// Пробежаться по массиву и добавить к code += 1
	fmt.Println(items[1].Code)
	fmt.Println(itemOne, itemTwo, itemThree)
}

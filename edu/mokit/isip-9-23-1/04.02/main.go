package main

import "fmt"

func main() {
	mapLit := map[string]int{"number": 100}
	mapLit["age"] = 10
	mapLit["number"] += 100
	fmt.Println("Способ «Литерал map»:", mapLit)
	fmt.Println("Search key 1:", mapLit["number"])

	value, exists := mapLit["number"]
	if exists {
		fmt.Println("Search key 2:", value)
	}
}

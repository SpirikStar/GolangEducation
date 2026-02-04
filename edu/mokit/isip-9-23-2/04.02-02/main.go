package main

import "fmt"

func main() {
	mapOne := map[string]int{"year": 2026}

	mapOne["number"] = 10
	delete(mapOne, "number")

	value, exists := mapOne["year"]
	if exists {
		fmt.Println("Read «Литерал»:", value)
	}
	fmt.Println("Способ «Литерал»:", mapOne)

	mapTwo := make(map[string]int)
	fmt.Println("Способ «make»:", mapTwo)

	for key, value := range mapOne {
		fmt.Printf("%s -> %d\n", key, value)
	}
	
	// TODO: Используйте for, fmt.Scan
	// TODO: 1. Запросите у пользователя 5 слов.
	// ! 2. Если в слове есть пробел, не сохранять!
	// TODO: 3. Сохраните в «Литерал» slice
	// TODO: 4. Создать «Литерал» map с результатом:
	// ? [hello:5 home:4 bg:2]
}

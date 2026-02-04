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

	mapMake := make(map[string]int)
	mapMake["age"] = 10
	delete(mapMake, "age")
	fmt.Println("Способ «make map»:", mapMake)
	fmt.Println("Len «make map»:", len(mapMake))

	for key, value := range mapLit {
		fmt.Printf("%s -> %d \n", key, value)
	}

	// ! Дедлайн: 13:05
	// TODO: Лекция
	// ? Чем отличает литерал от make  // 1.5
	// ? Что вы получите, если найти ключ в литерале // 1.0
	// ? Что означение запись []int{1, 2, 3} // 0.5
	for {
		var word string
		fmt.Scan(&word)
	}
	// TODO: Практика
	// Запросите у пользователя 2 слова //  0.5
	// Сохраните 2 слова в литерал slice //  0.5
	// Отобразить результат в литерале map: // 1.0
	// TODO: [home:4 hello:5]
}

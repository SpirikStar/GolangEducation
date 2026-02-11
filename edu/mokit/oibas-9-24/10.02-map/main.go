package main

import "fmt"

func main() {
	// ? Создание map: var переменная map[типКлюч]типЗначение

	var mapOne map[string]int
	fmt.Println(mapOne)

	mapLit := map[string]int{
		"yaer": 2026,
	}
	fmt.Println("Способ литерал: ", mapLit)
	fmt.Println("Способ литерал: ", mapLit["yaer"])

	mapMake := make(map[string]int)
	fmt.Println("Способ make: ", mapMake)

	mapMakeContainer := make(map[string]int, 5)
	mapMakeContainer["age"] = 7
	mapMakeContainer["age"] += 1
	mapMakeContainer["year"] = 2026

	delete(mapMakeContainer, "year")

	fmt.Println("Способ make: ", mapMakeContainer)
	fmt.Println("Способ make: ", mapMakeContainer["age"])
	value, exists := mapMakeContainer["year"]
	if exists {
		fmt.Println("Search key = value: ", value)
	}

	var userText string
	infosLit := map[string]int{"letters": 0}
	fmt.Print("Введите ссобщение: ")
	fmt.Scan(&userText)

	for _, value := range userText {
		if value != '.' {
			infosLit["letters"] += 1
		}
	}
	fmt.Println(infosLit)
}

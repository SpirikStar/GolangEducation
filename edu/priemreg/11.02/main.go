package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var varOne any = 100
	varOne = 10.2
	fmt.Println(varOne)

	sliceOne := []any{1, 2, 1.1, 3.1}
	sliceOne = append(sliceOne, 4.2)
	for _, value := range sliceOne {
		fmt.Printf("type: %T; value: %v\n", value, value)
	}

	matrixOne := [3][3]int{}
	matrixTwo := [3][3]int{{1, 2, 3}, {4, 5}}
	fmt.Println(matrixOne, matrixTwo)

	mapOne := map[string]any{
		"firstName": "Bobik",
		"year":      2026,
	}
	mapOne["age"] = 3
	delete(mapOne, "year")
	fmt.Println(mapOne)

	loc, _ := time.LoadLocation("Europe/Moscow")
	currentTime := time.Now().In(loc)
	fmt.Println("Date and time: ", currentTime)
	fmt.Println("Current hour: ", currentTime.Hour())
	fmt.Println("Current minute: ", currentTime.Minute())
	fmt.Println("Current second: ", currentTime.Second())
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Intn(100)) // от 0 до 100
	fmt.Println(r.Float64()) // от 0.1 до 1.0
	fmt.Println(r.Float32()) // от 0.1 до 1.0



}

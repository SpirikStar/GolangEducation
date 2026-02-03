package main

import "fmt"

func main() {
	fmt.Print("Проверка...")
	fmt.Println("Проверка...")
	fmt.Printf("Проверка... %d\n", 100)

	var float32One float32 = 20.1 // 6-7 / 4 байт
	var float64One float64 = 15.1 // 15-17 / 8 байт
	float64Two := 50.1            // default float64
	fmt.Println("float32One: ", float32One)
	fmt.Println("float64One: ", float64One)
	fmt.Println("float64Two: ", float64Two)

	var int32One int32 = 5    // -2 до 2 млрд.
	var int64One int64 = 10   // -9 до 9 млрд.
	var uint64One uint64 = 10 // 0 до 18 млрд.
	var int64Two int = 15     // auto system
	int64Three := 15          // default int64
	fmt.Println("int32One: ", int32One)
	fmt.Println("int64One: ", int64One)
	fmt.Println("int64Two: ", int64Two)
	fmt.Println("uint64One: ", uint64One)
	fmt.Println("int64Three: ", int64Three)

	var isBoolOne bool = true
	isBoolTwo := false
	fmt.Println("isBoolOne: ", isBoolOne)
	fmt.Println("isBoolTwo: ", isBoolTwo)

	var stringOne string = "Hello..."
	stringTwo := "Hello..."
	fmt.Println("stringOne: ", stringOne)
	fmt.Println("stringTwo: ", stringTwo)

	var a float32 = 5.12
	var b float64 = 1.19

	resOne := float64(a) + b
	resTwo := a + float32(b)
	fmt.Println("resOne: ", resOne)
	fmt.Println("resTwo: ", resTwo)

	origin := 0
	mask := &origin
	*mask++
	fmt.Println("origin:", origin)

	var number int = 500
	fmt.Println(number)

	var balance float32 = -5.9250000001
	fmt.Println(balance)

	var runeVar rune = 'Y'
	runeTwo := 'M'
	fmt.Println(runeVar, runeTwo)

	fmt.Printf(
		"Строка: %s\nЧисло: %d\n",
		"text", 100,
	)
	fmt.Printf(
		"Строка: %f\nЧисло: %.2f\n",
		10.123, 15.987,
	)
}

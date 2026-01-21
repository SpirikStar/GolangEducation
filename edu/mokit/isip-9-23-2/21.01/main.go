package main

import "fmt"

func main() {
	fmt.Println("Процесс...")

	stringOne := "text"
	var stringTwo string = "new text"
	fmt.Println("stringOne:", stringOne)
	fmt.Println("stringTwo:", stringTwo)

	intOne := 100            // auto
	var intTwo int = 200     // 32 bit or 64 bit
	var intThree int32 = 400 // -2 до 2 млрд.
	var intFour int64 = 600  // -9 до 9 млрд.
	fmt.Println(intOne, intTwo, intThree, intFour)

	isOne := true          // 1
	var isTwo bool = false // 0
	fmt.Println(isOne, isTwo)

	floatOne := 1.1              // auto; default 64
	var floatTwo float32 = 3.1   // .6-7
	var floatThree float64 = 5.1 // .15-16
	fmt.Println(floatOne, floatTwo, floatThree)

	a := 10
	b := 5.5
	var c float32 = 5.8
	resultOne := float64(a) + b
	resultTwo := a + int(b)
	resultThree := float32(b) + c
	fmt.Println(resultOne, resultTwo, resultThree)

	fInt := 1
	fFloat64 := 5.101517
	fBool := true
	fString := "text"

	message := fmt.Sprintf(
		"%d - %f / %.2f - %t - %s",
		fInt,
		fFloat64,
		fFloat64, 
		fBool, 
		fString,
	)
	fmt.Println(message)

	// TODO: Практическая. 
	varOne := 100
	varTwo := 3.1415926535897932384622433832795028841971
	// ! Использователь переменные для вывода 
	// TODO: 100.0 * 3.141592 = 314.15
}

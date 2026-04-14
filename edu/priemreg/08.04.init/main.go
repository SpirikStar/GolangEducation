package main

import (
	"log"
	"os"
)

var DEBAG *log.Logger

func init() {
	file, err := os.OpenFile("debag.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	DEBAG = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func factorial(n int) (int, any) {

	if n < 0 {
		return 0, "Недопустимое число. Должно быть больше 0"
	}

	result := 1
	for i := 1; i <= n; i += 1 {
		result *= i
	}
	return result, nil
}
func main() {
	// ? Пример
	_, errFile := os.Open("./content.txt")
	DEBAG.Println(errFile)

	// ? Пример
	fact, err := factorial(5)
	if err != nil {
		DEBAG.Println(err)
	} else {
		DEBAG.Println(fact)

	}
}

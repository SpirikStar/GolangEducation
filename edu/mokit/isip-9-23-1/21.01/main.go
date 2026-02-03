package main

import "fmt"

func main() {
	varExOne := 10
	varExTwo := varExOne // copy
	varExTwo++
	fmt.Println(varExOne, varExTwo)

	ex := &varExOne
	*ex++
	fmt.Println(ex, *ex)

	// ------------------
	elemRoot := 0
	var elem *int = &elemRoot
	*elem++
	fmt.Println(elem)
	// ------------------

	number := 90
	if number == 0 {
		fmt.Println("number == 0")
	} else if number > 50 && number < 80 {
		fmt.Println("number > 50 && number < 80")
	} else if number == 90 || number == 95 {
		fmt.Println("number == 90 || number == 95")
	} else {
		fmt.Println("Исключение")
	}

	balance := 200.5
	switch balance {
	case 200.5:
		fmt.Println(true)
	case 400.5:
		fmt.Println(false)
	default:
		fmt.Println(nil)
	}

	switch {
	case balance > 0.0 && balance < 100.0:
		fmt.Println("balance между 0.0 и 100.0")
	default:
		fmt.Println(nil)
	}

	userBalance := 100.63
	fmt.Println(userBalance)
}

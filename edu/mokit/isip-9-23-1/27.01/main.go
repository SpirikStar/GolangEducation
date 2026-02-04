package main

import "fmt"

func main() {
	count := 0
	for count < 5 {
		fmt.Println(count)
		count += 1
	}

	countTwo := 0
	for {
		if countTwo < 5 {
			countTwo += 1
		} else {
			break
		}
	}
	fmt.Println(countTwo)

	countThree := 0

loop:
	for {
		switch countThree {
		case 5:
			break loop
		default:
			countThree += 1
		}
	}
	fmt.Println(countThree)

	numberOne, numberTwo := 10, 20
	fmt.Println(numberOne, numberTwo)
	text := "1Bobik"
	for index, value := range text {
		fmt.Printf(
			"index: %d; value: %s\n", index, string(value),
		)
	}
}

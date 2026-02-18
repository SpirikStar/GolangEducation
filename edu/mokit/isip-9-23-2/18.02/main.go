package main

import "fmt"

type Person struct {
	FirstName string
	Age       int
	Active    bool
}

func main() {
	userOne := Person{"Bobik", 3, true}
	userOne.Age += 1
	userTwo := Person{"Ivan", 12, false}
	userTwo.Active = true
	fmt.Println(userOne, userTwo)
	fmt.Println(userOne.Age)

	var userThree Person
	userThree.FirstName = "Gleb"
	userThree.Age = 10
	fmt.Println(userThree)

	numOne := func() int {
		return 100
	}
	numTwo := func(a, b int) int {
		return a + b
	}
	func() {
		fmt.Println("Start func!")
	}()
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)
	fmt.Println(numOne(), numTwo(1, 2))
}

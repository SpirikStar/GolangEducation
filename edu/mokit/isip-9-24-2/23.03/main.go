package main

import "fmt"

type User struct {
	FirstName string
	Year      uint16
}

func main() {
	userOne := User{}
	userOne.FirstName = "Bobik"
	fmt.Println(userOne)

	userTwo := User{"Ivan", 2026}
	fmt.Println(userTwo)

	var userThree User = User{
		Year: 2025,
	}
	fmt.Println(userThree)

	// TODO: array := []type{...}
	users := []User{
		{"Ivan", 2024},
		{"Gleb", 2026},
	}
	fmt.Println(users)
}

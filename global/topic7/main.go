package main

import "fmt"

type User struct {
	Name string
	Age  int
	City string
}

func (user User) Greet() string {
	return fmt.Sprintf("%s - %d - %s", user.Name, user.Age, user.City)
}
func (user User) ReplaceAgeOne() {
	// Не может менять состояние структуры÷
	user.Age += 1
}
func (user *User) ReplaceAgeTwo() {
	user.Age += 1
}
func main() {
	userOne := User{"Bobik", 3, "Подольск"}
	userOne.ReplaceAgeOne()
	userOne.ReplaceAgeTwo()
	fmt.Println(userOne.Greet())
}

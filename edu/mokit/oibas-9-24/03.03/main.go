package main

import "fmt"

type User struct {
	Login string
	Code  int
}

func (u *User) GetUser() string {
	return fmt.Sprintf(
		"login: %s = %d\n", u.Login, u.Code,
	)
}
func (u *User) ReplaceCode(new_code int) {
	u.Code = new_code
}

type Contact struct {
	Email string
}

type Person struct {
	Contact
	FirstName string
	Age       int
}

func main() {
	user := User{"admin", 1234}
	fmt.Println(user.GetUser())
	user.ReplaceCode(4567)
	print(user)

	personOne := Person{
		Contact: Contact{
			Email: "gleb@ya.ru",
		},
		FirstName: "Gleb",
		Age:       3,
	}
	fmt.Println(personOne)
	fmt.Println(personOne.Email)
	fmt.Println(personOne.Age == 10)
	personOne.Age += 2
	fmt.Println(personOne.Age)
	personTwo := Person{}
	fmt.Println(personTwo)
}

// TODO: 1. Что такое метод?
// TODO: 2. Для чего используется struct?
// TODO: 3. Что такое наследование?
// TODO: 4. Дедлайн 11:15
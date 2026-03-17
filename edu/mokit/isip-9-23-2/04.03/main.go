package main

import "fmt"

type Auth struct {
	Login    string
	Password string
}

type User struct {
	Auth
	FirstName string
}

type ImageStyle struct {
	Width  int
	Height int
}
type Image struct {
	ImageStyle
	Src string
	Alt string
}

func main() {
	userOne := User{}
	userOne.Login = "test"
	fmt.Println(userOne, userOne.FirstName)

	userTwo := User{
		Auth: Auth{
			Login:    "gleb@ya.ru",
			Password: "1234",
		},
		FirstName: "Gleb",
	}
	fmt.Println(userTwo)

	images := []Image{}
	images = append(images, Image{Src: "http://photo.ru/1"})
	images = append(images, Image{Src: "http://photo.ru/2"})
	fmt.Println(images)

	// TODO: Создать 2 структуры и связать их
	// TODO: 1 - Видео
	// TODO: 2 - Настройки к видео

	// TODO: Создать массив, который будет хранить 
	// TODO: пользовательские видео

	// TODO: Добавление видео в массив должно 
	// TODO: происходить через функцию
}

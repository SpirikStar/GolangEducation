package main

import "fmt"

type User struct {
	FirstName string
	Balance   float64
	Status    bool
}
type Image struct {
	Src    string
	Alt    string
	Width  int
	Height int
}

func (image *Image) getSrc() {
	fmt.Println(image.Src)
}
func (user *User) upgradeBalance(elem float64) {
	user.Balance += elem
}
func main() {
	userOne := User{}
	fmt.Println(userOne)

	userTwo := User{"Bobik", 120.5, true}
	userTwo.Balance += 10.5
	fmt.Println(userTwo)
	fmt.Println(userTwo.FirstName)

	userThree := User{Balance: 250.2, FirstName: "Ivan", Status: false}
	userThree.upgradeBalance(100.0)
	fmt.Println(userThree)

	images := []Image{}
	images = append(
		images,
		Image{Src: "photo-1.png", Width: 100},
		Image{Src: "photo-2.png", Width: 100},
		Image{Src: "photo-3.png", Width: 100},
	)
	for _, value := range images {
		value.getSrc()
	}
}

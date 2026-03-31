package main

import "fmt"

type Item struct {
	Name string
}

func main() {
	item := struct {
		Name string
	}{"Bobik"}
	fmt.Println("struct (1):", item.Name)

	// Array: []type{...}
	items := []Item{
		{"User 1"},
		{"User 2"},
	}
	for _, elem := range items {
		fmt.Println(elem.Name)
	}

	users := []struct {
		FirstName string
		Yaer      uint16
	}{
		{"Bobik", 2026},
		{"Gleb", 2025},
	}
	fmt.Println(users)
}

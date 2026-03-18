package main

import "fmt"



func main() {
	user := struct {
		FirstName string
		Year      uint32
	}{
		FirstName: "Bobik",
		Year:      2026,
	}
	fmt.Println(user)

	nums := []struct {
		a int
		b int
	}{
		{a: 1, b: 2},
		{b: 1, a: 2},
		{2, 3},
	}
	fmt.Println(nums)
}

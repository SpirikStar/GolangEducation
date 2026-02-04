package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println(a, b, c)

	var nums [3]int
	fmt.Println(nums)

	numsTwo := [3]int{1, 20}
	numsTwo[2] = 30
	fmt.Println(numsTwo)
	fmt.Println(len(numsTwo))
	fmt.Println(numsTwo[0])

	for index, value := range numsTwo {
		fmt.Println(index, value)
	}
}

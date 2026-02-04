package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println(a, b, c)

	var nums [3]int
	fmt.Println(nums)

	numsTwo := [3]int{10, 20}
	numsTwo[0] += 40
	numsTwo[2] = 100
	fmt.Println(numsTwo)
	fmt.Println(len(numsTwo))
	fmt.Println(numsTwo[0])

	var matrix [2][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	fmt.Println(matrix[1])
	fmt.Println(matrix[0][1])
	fmt.Println(matrix)

	arrayOne := [2]int{1, 2}
	arrayTwo := [2]int{3, 4}
	fmt.Println(arrayOne == arrayTwo) // false

	elems := [4]int{100, 200, 300, 400}
	for index, value := range elems {
		fmt.Println(index, value)
	}
}

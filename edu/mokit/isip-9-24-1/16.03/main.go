package main

import "fmt"

func actionNumber(n int8) (bool, int8) {
	if n <= 0 {
		return false, 1
	}
	return true, n
}

func main() {
	status, res := actionNumber(-10)
	fmt.Println(status, res)

	func(elems ...int) {
		fmt.Println(elems)
	}(1, 2, 3)

	_, val := actionNumber(func(n int) int8 {
		if n < 255 {
			return int8(n)
		}
		return 0
	}(500))
	fmt.Println(val)

	result := func(val int) string {
		return fmt.Sprintf("Value: %d", val)
	}
	fmt.Println(result(1))
	fmt.Println(result(2))
}

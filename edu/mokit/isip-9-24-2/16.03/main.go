package main

import "fmt"

func cacl(a, b int) {
	fmt.Println(a * b)
}
func main() {
	func() {
		fmt.Println("Funcs (1)")
	}()
	cacl(func(valOne int) int {
		if valOne == 0 {
			return 1
		}
		return valOne
	}(0), 1)

	funcNeo := func() string {
		return "Func neo (2)"
	}
	fmt.Println(funcNeo())
	fmt.Println(funcNeo())
}

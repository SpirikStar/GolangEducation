package main

import "fmt"

func division(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("ERROR E101")
	}
	return a / b, nil
}
func main() {	
	fmt.Println(division(1, 0))
	fmt.Println(division(1, 2))
}

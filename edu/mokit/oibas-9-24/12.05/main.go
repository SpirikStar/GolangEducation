package main

import "log"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
func sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}
func attemptOperation(n int) int {
	if n <= 0 {
		return n
	}
	res := attemptOperation(n - 1)
	log.Println(res)
	return res + 1
}
func main() {
	log.Printf("5! = %d\n", factorial(5))
	log.Printf("sum %d = %d\n", 10, sum(10))

	attempt := 5
	attemptOperation(attempt)
}

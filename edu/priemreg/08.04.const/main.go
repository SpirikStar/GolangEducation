package main

const config = "dark"

func main() {
	const elem = 100
	//! elem += 1

	var varElem int = elem
	var varElemTwo float64 = elem
	varElemTwo++
	varElem++

	//! const n int = 10

	// ? () - группа
	const (
		one   = 1
		two   = 2
		three = 3
	)
}

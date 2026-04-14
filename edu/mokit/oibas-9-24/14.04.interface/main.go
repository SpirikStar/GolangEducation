package main

import "fmt"

type Elem struct {
	Name string
}
type NewElem struct {
	Name any
}

func (e *Elem) getElem() {
	fmt.Println("+", e.Name)
}

func (e *NewElem) getElem() {
	fmt.Println("-", e.Name)
}

type Operation interface {
	getElem()
}

func main() {
	var operation Operation
	operation = &Elem{Name: "test one"}
	operation.getElem()

	operation = &NewElem{Name: "test two"}
	operation.getElem()
}

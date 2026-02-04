// TODO: any, матрица, map
package main

import "fmt"

func main() {
	// [длинна]типДанных - массив
	var arrOne [2]int
	arrOne[1] = 5
	arrOne[1] += 5
	fmt.Println(arrOne)

	arrTwo := [5]int{10, 20}
	fmt.Println(arrTwo)

	var arrThree [5]int = [5]int{10, 20}
	fmt.Println(arrThree)

	for index, value := range arrThree {
		fmt.Printf("index: %d; value: %d \n", index, value)
	}

	// []типДанных - срез
	var sliceOne []float32
	sliceOne = append(sliceOne, 1.1)
	sliceOne = append(sliceOne, 1.3, 1.2, 1.4)

	var sliceOneCopy []float32
	for _, value := range sliceOne {
		if value == 1.2 {
			continue
		}
		sliceOneCopy = append(sliceOneCopy, value)
	}
	fmt.Println(sliceOneCopy)

	sliceTwo := []float32{1.0, 2.0, 3.0}
	sliceTwo = append(sliceTwo, 4.0)
	fmt.Println(sliceTwo)

	sliceMake := make([]int, 10)
	sliceMake = append(sliceMake, 1)
	fmt.Println(sliceMake)

	sliceMakeTwo := make([]int, 0, 2)
	sliceMakeTwo = append(sliceMakeTwo, 2, 3)
	fmt.Println(len(sliceMake), cap(sliceMakeTwo))

	// map[типДанныхКлюча]типДанныхЗначения
	infoUser := map[string]int{"number": 100, "yaer": 2026}
	infoUser["age"] = 10
	infoUser["age"] += 5
	fmt.Println(infoUser, infoUser["age"])
}

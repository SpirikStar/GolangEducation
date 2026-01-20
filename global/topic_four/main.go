package main

import "fmt"

func main() {
	mOne := make(map[string]int)
	fmt.Println(mOne)

	mTwo := map[string]int{
		"age_one": 10,
		"age_two": 20,
	}
	mTwo["age_three"] = 30
	mTwo["age_one"] += 5
	fmt.Println(mTwo)

	mThree := map[string]any{
		"name": "Alice",
		"age":  30,
	}
	fmt.Println(mThree)
	delete(mTwo, "age_two")
	value, exists := mTwo["age_two"]
	if exists {
		fmt.Println("age_two:", value)
	} else {
		fmt.Println("age_two does not exist")
	}

	for key, value := range mTwo {
		fmt.Printf("%s: %d\n", key, value)
	}

	
}

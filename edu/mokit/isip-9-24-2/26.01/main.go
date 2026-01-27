package main

import "fmt"

func main() {
	var float64One float64 = 20.5 // 15–17
	var float32One float32 = 5.5  // 6-7
	float64Two := 6.1             // default float64
	fmt.Println("float64One:", float64One)
	fmt.Println("float32One:", float32One)
	fmt.Println("float64Two:", float64Two)
	
	resOne := float32(float64One) + float32One
	resTwo := float64One + float64(float32One)
	fmt.Println("resOne:", resOne)
	fmt.Println("resTwo:", resTwo)

	var int32One int32 = 10 // - 2 до 2 млрд.
	var int64One int64 = 20 // -9 до 9 млрд.
	var int64Two uint64 = 20 // 0 до 18 млрд.
	var intOne int = 30     // auto system
	intTwo := 40            // auto system
	fmt.Println("int32One:", int32One)
	fmt.Println("int64One:", int64One)
	fmt.Println("int64Two:", int64Two)
	fmt.Println("intOne:", intOne)
	fmt.Println("intTwo:", intTwo)
	
	var isBoolOne bool = true
	isBoolTwo := false
	fmt.Println("isBoolOne:", isBoolOne)
	fmt.Println("isBoolTwo:", isBoolTwo)
	
	var stringOne string = "Hello"
	stringTwo := "Hello"
	fmt.Println("stringOne:", stringOne)
	fmt.Println("stringTwo:", stringTwo)


	var varBlock float32 = 1.00000178
	fmt.Println(varBlock)

	fmt.Print(100)
}

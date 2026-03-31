package main

import "fmt"

func main() {
	// {Message 1, Message 2}
	elems := make(chan string, 5)

	go func() {
		fmt.Println("go: send message 1")
		elems <- "Message 1"

		fmt.Println("go: send message 2")
		elems <- "Message 2"

		fmt.Println("go: send message 3")
		elems <- "Message 3"

		fmt.Println("go: send message ok!")
	}()

	fmt.Println("main: wait action user")
	fmt.Scanln()

	message1 := <-elems
	fmt.Println("main: response 1 =", message1)

	message2 := <-elems
	fmt.Println("main: response 2 =", message2)

	message3 := <-elems
	fmt.Println("main: response 3 =", message3)

	message4 := <-elems
	fmt.Println("main: response 4 =", message4)
}

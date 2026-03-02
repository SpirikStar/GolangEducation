package main

import (
	// "fmt"
	"os"
)

func main() {
	content := []byte("is Bobik?")
	os.WriteFile("info.txt", content, 0644)

	data, err := os.ReadFile("info.txt")
	if err != nil {
		return
	}
	os.Stdout.Write(data)

}

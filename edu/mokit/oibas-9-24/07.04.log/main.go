package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var DEBAG *log.Logger

func init() {
	file, err := os.OpenFile("debag.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	DEBAG = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func getImage(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		DEBAG.Println(err)
		return []byte{}
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		DEBAG.Println(response.StatusCode)
		return []byte{}
	}
	content, err := io.ReadAll(response.Body)
	if err != nil {
		DEBAG.Println(err)
		return []byte{}
	}
	return content
}

// TODO: https://loremflickr.com/
func main() {
	byteImage := getImage("https://loremflickr.com/800/600/cat")
	log.Println(byteImage)
}

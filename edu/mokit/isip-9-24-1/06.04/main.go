package main

import (
	"log"
	"net/http"
	"os"
)

var Debag *log.Logger

func init() {
	file, err := os.OpenFile("DEBAG.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	Debag = log.New(file, "DEBAG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	response, err := http.Get("https://clck.ru")
	if err != nil {
		Debag.Println(err)
	}
	Debag.Println(response.StatusCode, response.Header)
	// TODO: https://metanit.com/go/tutorial/9.5.php
	// TODO: GET https://clck.ru
	// TODO: Получить статус и headers
}

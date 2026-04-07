package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var Debag *log.Logger

func init() {
	file, err := os.OpenFile("debag.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = loc
	Debag = log.New(file, "DEBAG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	endpoint := "https://clck.ru/--"
	url := "https://ya.ru"
	response, err := http.Get(endpoint + "?url="+url)
	if err != nil {
		Debag.Println(err)
		return
	}
	fmt.Println(response.Body)
}

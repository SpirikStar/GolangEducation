package main

import (
	"log"
	"os"
)

var debag *log.Logger

func init() {
	file, _ := os.OpenFile(
		"debag.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	debag = log.New(
		file,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}

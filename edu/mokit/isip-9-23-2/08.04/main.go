package main

import (
	"log"
	"os"
)

var DEBAG *log.Logger

func init() {
	file, err := os.OpenFile("debag.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	DEBAG = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// TODO: Изучить всю 10 главу. 
	// TODO: Используйте текущий код и добавляйте 
	// TODO: примеры представленные в документации  
	// TODO: https://metanit.com/go/tutorial/11.1.php
	// TODO: Логируйте результат!
	DEBAG.Println(200)
}

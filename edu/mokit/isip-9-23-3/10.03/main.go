// ! go mod init ws
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", echoHandle)
	http.HandleFunc("/", indexPage)
	log.Println(http.ListenAndServe(":8080", nil))
}

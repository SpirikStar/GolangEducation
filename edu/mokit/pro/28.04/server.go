package main

import "net/http"

var clients []string

func wsHandle(w http.ResponseWriter, r *http.Request) {

}

func pageHandle(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("GET /", pageHandle)
	http.HandleFunc("/ws", wsHandle)
}

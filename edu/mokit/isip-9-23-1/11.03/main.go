package main

import (
	"html/template"
	"net/http"
)

type Info struct {
	Number uint16
	Text   string
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html, _ := template.ParseFiles("index.html")
	html.Execute(w, Info{
		Number: 10,
		Text:   "New text",
	})
}

func main() {
	http.HandleFunc("/", handleHome)
	http.ListenAndServe(":8080", nil)
}

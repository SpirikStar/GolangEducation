package main

import (
	"net/http"
)

type Image struct {
	Src    string
	Alt    string
	Width  int
	Height int
}

var images = []Image{}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not method post", 405)
	}
}

func getImages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not method get", 405)
	}
}

func main() {

	http.HandleFunc("/upload-image/", uploadImage)
	http.HandleFunc("/", getImages)
	
	http.ListenAndServe(":8000", nil)
}

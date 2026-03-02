package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", search)
}

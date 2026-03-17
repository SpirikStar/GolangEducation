package main

import (
	"fmt"
	"log"
	"net/http"
)
// TODO: https://pkg.go.dev/net/http#Request

func handleStart(w http.ResponseWriter, r *http.Request) {
	log.Println("host:", r.Host)
	log.Println("method:", r.Method)
	log.Println("client:", r.UserAgent())
	fmt.Fprintln(w, "site...")
}

func main() {
	http.HandleFunc("/", handleStart)

	log.Println("Start http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

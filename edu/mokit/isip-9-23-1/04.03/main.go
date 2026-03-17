// ! https://metanit.com/go/tutorial/7.1.php
package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"
)

func handleStart(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	log.Println(r.URL)
	log.Println(r.UserAgent())
	fmt.Fprintln(w, "Site...")
}
func main() {
	log.Println("host: http://localhost:8000/")

	http.HandleFunc("/", handleStart)
	http.ListenAndServe(":8000", nil)
}

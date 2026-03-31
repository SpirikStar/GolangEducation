package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "0f5838b93d044ec582b1866f7ca627bf.png")
	})
	http.ListenAndServe(":8080", nil)
}

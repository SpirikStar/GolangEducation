// go run main.go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", startHandle)
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func startHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	value := r.URL.Query().Get("value")

	file, _ := os.Open("words.txt")
	defer file.Close()
	
	scanner := bufio.NewScanner(file)

	result_words := []string{}
	for scanner.Scan() {
		word_file := strings.ToLower(scanner.Text())
		if strings.HasPrefix(word_file, value) {
			result_words = append(result_words, word_file)
		}
	}

	response := struct {
		Success bool     `json:"success"`
		Words   []string `json:"words"`
	}{Success: len(result_words) > 0, Words: result_words}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

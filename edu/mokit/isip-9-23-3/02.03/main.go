package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status bool `json:"status"`
}

func main() {
	http.HandleFunc("/", handleStart)
	http.ListenAndServe(":8000", nil)
}
func handleStart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	log.Println(r.RemoteAddr)
	log.Println(r.UserAgent())
	// TODO: Создать базовую защиту DDOS-атаки
	// TODO: Не больше 10 запросов в минуту
	data := Response{
		Status: true,
	}
	json.NewEncoder(w).Encode(data)
}

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	City      string `json:"city"`
	Age       uint   `json:"age"`
	CreatedAt string `json:"created_at"`
}

var db *sql.DB
// TODO: http://localhost:8080/users
// TODO: GET /users?age__lt
// TODO: GET /users?age__gt
// TODO: GET /users?age__lt={}&age__gt={}
func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT * FROM users")
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		rows.Scan(
			&user.ID, &user.FirstName, &user.LastName,
			&user.Email, &user.City, &user.Age,
			&user.CreatedAt,
		)
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func main() {
	var err error

	db, err = sql.Open("sqlite3", "./messenger.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", getUsers)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}

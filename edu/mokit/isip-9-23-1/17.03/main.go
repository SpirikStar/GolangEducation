package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	City      string
	Age       uint
	CreatedAt string
}

type Message struct {
	ID         uint
	SenderID   uint
	ReceiverID uint
	Text       string
	SentAt     string
	IsRead     bool
}

func openDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalln("Ошибка:", err)
	}
	return db
}
func main() {
	db := openDB("messenger.db")
	defer db.Close()

	var user User
	db.QueryRow("SELECT * FROM users WHERE id = ?", 1).Scan(
		&user.ID, &user.FirstName, &user.LastName,
		&user.Email, &user.City, &user.Age,
		&user.CreatedAt,
	)
	log.Println(user)
}

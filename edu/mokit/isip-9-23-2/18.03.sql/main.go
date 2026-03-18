package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
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

func connectDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return db
}

func getUsersAll(db *sql.DB) {
	var user User
	db.Query("SELECT * FROM user")
}
func main() {
	db := connectDB("messenger.db")
	defer db.Close()

	getUsersAll(db)
}

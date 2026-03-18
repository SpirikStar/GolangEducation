package main

import (
	"database/sql"
	"fmt"
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
	var users []User
	rows, _ := db.Query("SELECT * FROM users")

	for rows.Next() {
		var user User
		rows.Scan(
			&user.ID, &user.FirstName,
			&user.LastName, &user.Email,
			&user.City, &user.Age,
			&user.CreatedAt,
		)
		users = append(users, user)
	}
	fmt.Println(users)
}
// TODO: Создать функция поиска по id
// TODO: Создать функция поиска 
// TODO: пользователей > age
func main() {
	db := connectDB("messenger.db")
	defer db.Close()

	getUsersAll(db)
}

// TODO: go mod init sql-lesson
// TODO: go get modernc.org/sqlite
// TODO: go run .
package main

import (
	"database/sql"
	"log"
	"math/rand/v2"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite", "./base.db")
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	defer DB.Close()
	DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			first_name TEXT    NOT NULL,
			year       INTEGER DEFAULT 2026
		)
	`)
	yaer := rand.IntN(2026)
	DB.Exec(`
		INSERT INTO users (first_name, year)
		VALUES (?, ?)
	`, "Bobik", yaer)
	
	DB.Exec(`
		SELECT * FROM users
	`)

	// TODO: 1. Сделать генерацию для `year`
	// TODO: 2. Выполнить SELECT запрос
}

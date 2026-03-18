// TODO: go mod init lesson
package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func getUsersAll() {}
func connectDB(path string) {
	sql.Open("sqlite3", path)
}
func main() {
	connectDB("messenger.db")
}

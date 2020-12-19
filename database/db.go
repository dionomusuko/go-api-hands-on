package database

import (
	"database/sql"
	"os"
)

// TODO DBへの接続処理を書く
func NewSqLHandler() *sql.DB {
	dbURI := "user:password@tcp(db:3306)/go_db"
	conn, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Printf("connected error: %v", dbURI)
		os.Exit(1)
	}
}

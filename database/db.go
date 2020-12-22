package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var Conn *sql.DB

func init() {
	dbURI := "user:password@tcp(db:3306)/go_db"
	conn, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Printf("connected error: %v", dbURI)
		os.Exit(1)
	}
	Conn = conn
}

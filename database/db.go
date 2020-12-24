package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {

	dbURI := "user:password@tcp(db:3306)/go_db"
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Printf("connected error: %v", dbURI)
		os.Exit(1)
	}
	DB = db
}

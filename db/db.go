package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

	connStr := "host=localhost port=5432 user=postgres password=Rehna@1480 dbname=crm sslmode=disable"

	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	log.Println("✅ PostgreSQL Connected Successfully!")
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	if user == "" || password == "" || host == "" || dbname == "" {
		log.Fatal("Missing required environment variables: DB_USERNAME, DB_PASSWORD, DB_HOST, DB_NAME")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, dbname)

	var db *sql.DB
	var err error
	for i := 0; i < 30; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("Waiting for MySQL to be ready (%d/30)...", i+1)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatalf("Error connecting to database after retries: %v", err)
	}

	// Create table if not exists
	createTable := `CREATE TABLE IF NOT EXISTS items (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT
	)`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	return db
}

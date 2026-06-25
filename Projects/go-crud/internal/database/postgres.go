package database

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	_"github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error){
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Opening db: %w", err)
	}

	log.Println("Database connected successfully")

	if err := db.Ping();
	err != nil {
		return nil, fmt.Errorf("Connecting to db: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetConnMaxIdleTime(25)

	return db, nil
}
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

// NewConnection returns a new database connection instance
func NewConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	conn := fmt.Sprintf("host=%s user=%s port=5432 dbname=%s password=%s sslmode=disable", host, user, dbName, password)
	fmt.Println(conn)
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

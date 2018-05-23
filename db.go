package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// nolint: gotype
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// DB is the DB that will performs all operation
type DB struct {
	DB *gorm.DB
}

// NewDB returns a new DB connection
func newDB() (*DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("db_user")
	dbPass := os.Getenv("db_pass")
	// connect to the example db, create if it doesn't exist
	db, err := gorm.Open("mysql", dbUser+":"+dbPass+"@/learngraphql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

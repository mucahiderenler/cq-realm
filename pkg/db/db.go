package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Config holds the database configuration
// TODO: Load this from out .env config file
type Config struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
	SSLMode  string
}

// ProvideDB initializes the database connection
func ProvideDB() (*sqlx.DB, error) {
	config := Config{
		User:     "postgres",
		Password: "password",
		DBName:   "cqrealm",
		Host:     "127.0.0.1",
		Port:     "5432",
		SSLMode:  "disable",
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		config.User, config.Password, config.DBName, config.Host, config.Port, config.SSLMode)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connection established")
	return db, nil
}

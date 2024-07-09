package services

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	// Load the environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	// Retrieve the database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// Connect to the database
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

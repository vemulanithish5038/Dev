package db

import (
	"database/sql"
	"ecommerce-crud-golang/config"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection
// Singleton pattern is used to ensure only one connection is created
// The connection is stored in the DB variable

func InitDB() {
	conf := config.AppConfig
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// TODO
	// Add more error & exception handling
	// The micoservice should not panic
	// It should reconnect when the DB is live again

	DB = db
}

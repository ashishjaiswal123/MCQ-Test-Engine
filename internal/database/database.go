package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB *gorm.DB
)

// Connect initializes the database connection
func Connect() error {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %s", err))
	}

	err = DB.DB().Ping()
	if err != nil {
		return err
	}

	DB.AutoMigrate(&User{})
	return nil
}

// Close closes the database connection
func Close() {
	if err := DB.Close(); err != nil {
		fmt.Printf("Error closing database connection: %s", err)
	}
}

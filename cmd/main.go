package main

import (
	"fmt"
	"test-engine/internal/database"
	"test-engine/internal/utils"
)

func main() {
	fmt.Println("Welcome to the Multiple Choice Test!")

	// Initialize database connection
	if err := database.Connect(); err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		return
	}
	defer database.Close()

	router := SetupRoutes()

	port := utils.GetEnv("PORT", "8080")
	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"test-engine/internal/testengine"
	"time"

	"github.com/gin-gonic/gin"
)

// StartTest is the handler function to start the test
func StartTest(c *gin.Context) {
	// Load the test from a JSON file
	test, err := testengine.LoadTestFromFile("example/example_test.json")
	if err != nil {
		fmt.Printf("Error loading test: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	testEngine := testengine.TestEngine{
		Test:        test,
		UserScore:   0,
		CurrentTime: time.Now(),
	}

	testEngine.RunTestEngine(c)
}

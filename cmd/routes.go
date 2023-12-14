package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Route to start the test
	router.POST("/start-test", StartTest)

	return router
}

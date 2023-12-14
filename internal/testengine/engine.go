package testengine

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"test-engine/internal/database"
	"test-engine/internal/model"

	"github.com/gin-gonic/gin"
)

// User struct to represent user details
type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	StartTime time.Time
}

// TestEngine struct represents the test engine
type TestEngine struct {
	UserScore   int
	CurrentTime time.Time
	Test        model.Test
}

// RunTestEngine runs the test engine
func (engine *TestEngine) RunTestEngine(c *gin.Context) {
	fmt.Println("Welcome to the Multiple Choice Test!")

	engine.ShuffleQuestions()

	user := engine.GetUserDetails(c)

	engine.CurrentTime = time.Now()

	for _, question := range engine.Test.Questions {
		engine.DisplayQuestion(&question)
		userAnswer := GetUserAnswer(c)
		question.UserAttempt = userAnswer
		engine.ValidateAnswer(&question)
	}

	engine.CurrentTime = time.Now()

	// Calculate the time taken
	duration := engine.CurrentTime.Sub(user.StartTime)
	minutes := int(duration / time.Minute)
	seconds := int(duration % time.Minute / time.Second)
	fmt.Printf("Test completed in %d minute %d second! You answered %d out of %d questions correctly.\n", minutes, seconds, engine.UserScore, len(engine.Test.Questions))

	// Store user details and score in the database
	dbUser := model.User{
		Name:      user.Name,
		Email:     user.Email,
		Score:     engine.UserScore,
		StartTime: user.StartTime,
		EndTime:   engine.CurrentTime,
	}

	if err := database.DB.Create(&dbUser).Error; err != nil {
		fmt.Printf("Error storing user details in the database: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test completed successfully!", "score": engine.UserScore})
}

// GetUserDetails gets user details before the test starts
func (engine *TestEngine) GetUserDetails(c *gin.Context) User {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Printf("Error getting user details: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		c.Abort()
		return user
	}

	user.StartTime = time.Now()
	return user
}

// DisplayQuestion displays the question and its options to the user
func (engine *TestEngine) DisplayQuestion(question *model.Question) {
	fmt.Printf("Question: %s\n", question.Text)
	for i, option := range question.Options {
		fmt.Printf("%c. %s ", 'A'+i, option)
	}
	fmt.Printf("\nYour answer: ")
}

// GetUserAnswer prompts the user for an answer and returns it
func GetUserAnswer(c *gin.Context) string {
	var userAnswer string
	fmt.Scanln(&userAnswer)
	userAnswer = strings.TrimSpace(userAnswer)
	return userAnswer
}
